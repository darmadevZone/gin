package repositories

import (
	"errors"
	"gin-market/mock/models"

	"gorm.io/gorm"
)

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
	Create(newItem models.Item) (*models.Item, error)
	Update(updateItem models.Item) (*models.Item, error)
	Delete(itemId uint) error
}

// Memoryの設定
type ItemMemoryRepository struct {
	items []models.Item
}

// Delete implements IItemRepository.
func (r *ItemMemoryRepository) Delete(itemId uint) error {
	for i, v := range r.items {
		if v.ID == itemId {
			r.items = append(r.items[:i], r.items[i+1:]...)
			return nil
		}
	}
	return errors.New("item not fount")
}

// Update implements IItemRepository.
func (r *ItemMemoryRepository) Update(updateItem models.Item) (*models.Item, error) {
	for i, v := range r.items {
		if v.ID == updateItem.ID {
			r.items[i] = updateItem
			return &r.items[i], nil
		}
	}
	return nil, errors.New("unexpected error")
}

// Create implements IItemRepository.
func (r *ItemMemoryRepository) Create(newItem models.Item) (*models.Item, error) {
	newItem.ID = uint(len(r.items) + 1)
	r.items = append(r.items, newItem)
	return &newItem, errors.New("Itemが作成されませんでした")
}

// FindById implements IItemRepository.
func (r *ItemMemoryRepository) FindById(itemId uint) (*models.Item, error) {
	for _, v := range r.items {
		if v.ID == itemId {
			return &v, nil
		}
	}
	return nil, errors.New("item not found")
}

func (r *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}

func NewItemMemoryRepository(items []models.Item) IItemRepository {
	return &ItemMemoryRepository{items}
}

// DB 設定について定義
type ItemRepository struct {
	db *gorm.DB
}

// Create implements IItemRepository.
func (i *ItemRepository) Create(newItem models.Item) (*models.Item, error) {
	result := i.db.Create(&newItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newItem, nil
}

// Delete implements IItemRepository.
func (i *ItemRepository) Delete(itemId uint) error {
	delteItem, err := i.FindById(itemId)
	if err != nil {
		return err
	}

	result := i.db.Delete(&delteItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FindAll implements IItemRepository.
func (i *ItemRepository) FindAll() (*[]models.Item, error) {
	var items []models.Item
	result := i.db.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return &items, nil
}

// FindById implements IItemRepository.
func (i *ItemRepository) FindById(itemId uint) (*models.Item, error) {
	var item models.Item
	result := i.db.First(&item, itemId)
	if result.Error != nil {
		if result.Error.Error() != "record not found" {
			return nil, errors.New("item not found")
		}
		return nil, result.Error
	}
	return &item, nil
}

// Update implements IItemRepository.
func (i *ItemRepository) Update(updateItem models.Item) (*models.Item, error) {
	result := i.db.Save(&updateItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updateItem, nil
}

func NewItemRepository(db *gorm.DB) IItemRepository {
	return &ItemRepository{db: db}
}
