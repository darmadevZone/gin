package main

import (
	"gin-market/mock/controllers"
	"gin-market/mock/infra"
	"gin-market/mock/repositories"
	"gin-market/mock/services"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	r := gin.Default()
	r.GET("/items", itemController.FindAll)
	r.GET("/items/:id", itemController.FindById)
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
