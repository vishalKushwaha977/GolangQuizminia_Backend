package repositories

import (
	"go-backend/internal/models"
	"gorm.io/gorm"
 )

 func GetAllQuestions(db *gorm.DB)([]models.Question,error){
    var questions []models.Question

	err := db.Table("questions").Find(&questions).Error
	if err != nil {
		return nil, err
	}

	return questions, nil

 }