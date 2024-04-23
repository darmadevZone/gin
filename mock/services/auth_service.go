package services

import (
	"fmt"
	"gin-market/mock/models"
	"gin-market/mock/repositories"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type ISignup interface {
	Signup(email string, password string) error
}
type ISignout interface {
	Signout(email string, password string) error
}
type GoogleSignin interface {
	ISignup
	ISignout
}

type IAuthService interface {
	ISignup
	ISignout
	Login(email string, password string) (*string, error)
	GetUserFromToken(tokenString string) (*models.User, error)
}

type AuthService struct {
	repositories repositories.IAuthRepository
}

// GetUserFromToken implements IAuthService.
// JWT tokenのデコード
func (a *AuthService) GetUserFromToken(tokenString string) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	var user *models.User
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return nil, jwt.ErrTokenExpired
		}
		user, err = a.repositories.FindUser(claims["email"].(string))
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}

// Login implements IAuthService.
func (a *AuthService) Login(email string, password string) (*string, error) {
	foundUser, err := a.repositories.FindUser(email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	token, err := CreateToken(foundUser.ID, foundUser.Email)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// Signout implements IAuthService.
func (a *AuthService) Signout(email string, password string) error {
	panic("")
}

// Signup implements IAuthService.
func (a *AuthService) Signup(email string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user := models.User{
		Email:    email,
		Password: string(hashedPassword),
	}
	return a.repositories.CreateUser(user)
}

func NewAuthService(repository repositories.IAuthRepository) IAuthService {
	return &AuthService{repositories: repository}
}

func CreateToken(userId uint, email string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   userId,
		"email": email,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}
