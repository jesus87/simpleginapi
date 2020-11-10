package http

import (
	"github.com/banwire/api-exam/handlers"
	"github.com/gin-gonic/gin"
)

// Router define all routes http
func Router(router *gin.Engine) {

	v1 := router.Group("/v1")                                                 
	{
		v1.POST("/merchant", handlers.CreateMerchantHandler)
		v1.GET("/merchant", handlers.GetMerchantHandler)
		v1.PUT("/merchant", handlers.EditMerchantHandler)                                                                                                                                     
	}
}
