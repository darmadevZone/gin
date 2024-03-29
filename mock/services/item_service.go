package services

import (
	"gin-market/mock/dto"
	"gin-market/mock/models"
	"gin-market/mock/repositories"
)

type IItemService interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
	Create(createItemInput dto.CreateItemInput) (*models.Item, error)
	Update(itemId uint, updateItemInput dto.UpdateItemInput) (*models.Item, error)
	Delete(itemId uint) error
}

// Service -> IRepository
type ItemService struct {
	repositories repositories.IItemRepository
}

// Delete implements IItemService.
func (s *ItemService) Delete(itemId uint) error {
	return s.repositories.Delete(itemId)
}

// Update implements IItemService.
func (s *ItemService) Update(itemId uint, updateItemInput dto.UpdateItemInput) (*models.Item, error) {
	targetItem, err := s.repositories.FindById(itemId)
	if err != nil {
		return nil, err
	}

	if updateItemInput.Name != nil {
		targetItem.Name = *updateItemInput.Name
	}
	if updateItemInput.Price != nil {
		targetItem.Price = *updateItemInput.Price
	}
	if updateItemInput.Description != nil {
		targetItem.Description = *updateItemInput.Description
	}
	if updateItemInput.SoldOut != nil {
		targetItem.SoldOut = *updateItemInput.SoldOut
	}
	return s.repositories.Update(*targetItem)
}

// Create implements IItemService.
func (s *ItemService) Create(createItemInput dto.CreateItemInput) (*models.Item, error) {
	newItem := models.Item{
		Name:        createItemInput.Name,
		Price:       uint(createItemInput.Price),
		Description: createItemInput.Description,
		SoldOut:     false,
	}
	return s.repositories.Create(newItem)
}

// FindById implements IItemService.
func (s *ItemService) FindById(itemId uint) (*models.Item, error) {
	return s.repositories.FindById(itemId)
}

func NewItemService(repository repositories.IItemRepository) IItemService {
	return &ItemService{repositories: repository}
}

func (s *ItemService) FindAll() (*[]models.Item, error) {
	return s.repositories.FindAll()
}
