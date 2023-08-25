package server

import (
	"github.com/Feruz666/serve/gateway/client"
	"github.com/Feruz666/serve/gateway/config"
	"github.com/Feruz666/serve/gateway/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter(cfg *config.Config) *gin.Engine {
	router := gin.New()

	v1 := router.Group("v0.0.0")
	{

		controllers := controllers.NewControllers(client.NewAccountClient(cfg))
		accountsGroup := v1.Group("accounts")
		{
			accountsGroup.POST("/auth", controllers.Register)
		}
	}

	return router
}
