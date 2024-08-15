package routes

import (
	"user-management/internal/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")

	api.POST("/users", handlers.CreateUser)
	api.POST("/users/generateotp", handlers.GenerateOTP)
	api.POST("/users/verifyotp", handlers.VerifyOTP)
}
