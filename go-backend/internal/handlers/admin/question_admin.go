package admin
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go-backend/internal/service"
)

func GetAllQuestionsAdmin(c *gin.Context){
	questions , err := service.GetAllQuestions()
    if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, questions)
}