package routes

import (
	"github.com/cheildo/meme_coin_api/api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/meme_coins/api")
	{
		api.POST("/", controllers.CreateMemeCoin)
		api.GET("/:id", controllers.GetMemeCoin)
		api.PUT("/:id", controllers.UpdateDescription)
		api.DELETE("/:id", controllers.DeleteMemeCoin)
		api.POST("/:id/poke", controllers.PokeMemeCoin)
	}

	return router
}
