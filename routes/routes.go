package routes

import (
	"github.com/gin-gonic/gin"

	"emailchaser.com/backend-go/controllers"
)

func SetupRoutes(router *gin.Engine) {
	leadRoutes := router.Group("/lead")
	leadRoutes.POST("", controllers.CreateLead)
	leadRoutes.GET("/:id", controllers.GetLead)
	leadRoutes.DELETE("/:id", controllers.DeleteLead)
	leadRoutes.PUT("/:id", controllers.UpdateLead)
}
