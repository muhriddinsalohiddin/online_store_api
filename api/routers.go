package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/muhriddinsalohiddin/online_store_api/api/docs" // swag
	v1 "github.com/muhriddinsalohiddin/online_store_api/api/handlers/v1"
	"github.com/muhriddinsalohiddin/online_store_api/config"
	"github.com/muhriddinsalohiddin/online_store_api/pkg/logger"
	"github.com/muhriddinsalohiddin/online_store_api/services"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

// New ...
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	api := router.Group("/v1")
	// Books
	api.POST("/books", handlerV1.CreateBook)
	api.GET("/books/:id", handlerV1.GetBookById)
	api.PUT("/books/:id", handlerV1.UpdateBook)
	api.DELETE("books/:id", handlerV1.DeleteBook)
	api.GET("/books", handlerV1.ListBooks)
	// Categories
	api.POST("/categories", handlerV1.CreateCategory)
	api.GET("/categories/:id", handlerV1.GetCategoryById)
	api.PUT("/categories/:id", handlerV1.UpdateCategory)
	api.DELETE("categories/:id", handlerV1.DeleteCategoryById)
	api.GET("/categories", handlerV1.ListCategories)
	// Authors
	api.POST("/authors", handlerV1.CreateAuthor)
	api.GET("/authors/:id", handlerV1.GetAuthor)
	api.PUT("/authors/:id", handlerV1.UpdateAuthor)
	api.DELETE("authors/:id", handlerV1.DeleteAuthor)
	api.GET("/authors", handlerV1.ListAuthors)
	// Orders
	api.POST("/orders", handlerV1.CreateOrder)
	api.GET("/orders/:id", handlerV1.GetOrderById)
	api.PUT("/orders/:id", handlerV1.UpdateOrder)
	api.DELETE("orders/:id", handlerV1.DeleteOrder)
	api.GET("/orders", handlerV1.ListOrders)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
