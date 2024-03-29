package main

import (
	"gin-market/mock/controllers"
	"gin-market/mock/models"
	"gin-market/mock/repositories"
	"gin-market/mock/services"

	"github.com/gin-gonic/gin"
)

func main() {
	items := []models.Item{
		{ID: 1, Name: "item1", Price: 100, Description: "desc1", SoldOut: false},
		{ID: 2, Name: "item2", Price: 200, Description: "desc2", SoldOut: false},
		{ID: 3, Name: "item3", Price: 300, Description: "desc3", SoldOut: false},
	}

	itemRepository := repositories.NewItemMemoryRepository(items)
	ItemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(ItemService)

	r := gin.Default()
	r.GET("/items", itemController.FindAll)
	r.GET("/item/:id", itemController.FindById)
	// r.GET("/json", func(ctx *gin.Context) {
	// 	repository.JsonPlaceHolderRepository{}.JsonPlaceHolderDataFindAll()
	// })
	// 作成
	r.POST("/items", itemController.Create)
	// 更新
	r.PUT("/items/:id", itemController.Update)
	//削除
	r.DELETE("/items/:id", itemController.Delete)

	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080
}

