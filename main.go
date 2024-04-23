package main

import (
	"gin-market/mock/controllers"
	"gin-market/mock/infra"
	"gin-market/mock/middlewares"
	"gin-market/mock/repositories"
	"gin-market/mock/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	//item
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	//auth
	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	r := gin.Default()
	r.Use(cors.Default())
	/* item */
	itemRouter := r.Group("/items")
	itemRouterWithAuth := r.Group("/items", middlewares.AuthMiddleware(authService))

	// 検索
	itemRouter.GET("", itemController.FindAll)
	itemRouterWithAuth.GET("/:id", itemController.FindById)
	// 作成
	itemRouterWithAuth.POST("", itemController.Create)
	// 更新
	itemRouterWithAuth.PUT("/:id", itemController.Update)
	//削除
	itemRouterWithAuth.DELETE("/:id", itemController.Delete)

	/* auth */
	authRouter := r.Group("/auth")
	// サインアップ
	authRouter.POST("/signup", authController.Signup)
	// ログイン
	authRouter.POST("/login", authController.Login)

	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080
}

// func httpExample() {
// 	h1 := func(w http.ResponseWriter, _ *http.Request) {
// 		io.WriteString(w, "Hello from a HandleFunc #1!\n")
// 	}
// 	h2 := func(w http.ResponseWriter, _ *http.Request) {
// 		io.WriteString(w, "Hello from a HandleFunc #2!\n")
// 		http.Redirect(w, nil, "/404", http.StatusMovedPermanently)
// 	}
// 	notFound := http.NotFoundHandler()
// 	client := http.Client{}
// 	url, _ := url.Parse("")
// 	request := http.Request{URL: url}
// 	c, _ := client.Do(&request)
// 	fmt.Print(c)

// 	http.HandleFunc("/", h1)
// 	http.HandleFunc("/endpoint", h2)
// 	http.Handle("/404", notFound)

// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }
