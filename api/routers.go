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
	api.POST("/books", handlerV1.CreateBook)
	api.GET("/books/:id", handlerV1.GetBookById)
	api.PUT("/books/:id", handlerV1.UpdateBook)
	api.DELETE("books/:id", handlerV1.DeleteBook)
	api.GET("/books", handlerV1.ListBooks)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}