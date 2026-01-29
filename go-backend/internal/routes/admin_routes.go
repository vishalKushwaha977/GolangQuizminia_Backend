package routes

import (
	"go-backend/internal/handlers/admin"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine) {
	adminRoute := r.Group("/admin")
	{
		adminRoute.GET("/questions", admin.GetAllQuestionsAdmin)
		
	}
}
