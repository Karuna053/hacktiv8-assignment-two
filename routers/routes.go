package routers

import (
	"rest-api/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiRouter := router.Group("/api")
	{
		apiRouter.POST("/create", controllers.CreateOrder)
		// apiRouter.GET("/", controllers.GetAllData)
		// apiRouter.PUT("/update", controllers.UpdateDataOrderAndItem)
		// apiRouter.DELETE("/delete", controllers.DeleteDataOrderAndItem)
	}

	return router
}
