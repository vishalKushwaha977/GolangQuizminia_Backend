package service
import (
   "go-backend/internal/db"
   "go-backend/internal/models"
   "go-backend/internal/repositories"
          
)

func GetAllQuestions()([]models.Question,error){
	
	return repositories.GetAllQuestions(db.DB)
}