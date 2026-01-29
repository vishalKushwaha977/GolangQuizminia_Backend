package models

import "time"

type Question struct {
	
	QuestionsNo int64  `json:"questions_no"`
	Question    string `json:"question"`

	Option1       string   `json:"option1"`
	Option1Weight *float64 `json:"option1_weight"`

	Option2       string   `json:"option2"`
	Option2Weight *float64 `json:"option2_weight"`

	Option3       string   `json:"option3"`
	Option3Weight *float64 `json:"option3_weight"`

	Option4       string   `json:"option4"`
	Option4Weight *float64 `json:"option4_weight"`

	Option5       *string  `json:"option5"`
	Option5Weight *float64 `json:"option5_weight"`

	Image1       []byte   `json:"image1,omitempty"`
	Image1Weight *float64 `json:"image1_weight"`

	Image2       []byte   `json:"image2,omitempty"`
	Image2Weight *float64 `json:"image2_weight"`

	Image3       []byte   `json:"image3,omitempty"`
	Image3Weight *float64 `json:"image3_weight"`

	Image4       []byte   `json:"image4,omitempty"`
	Image4Weight *float64 `json:"image4_weight"`

	Answer    string     `json:"answer"`
	UpdatedAt *time.Time `json:"updated_time_stamp"`

	CategoryID *int64 `json:"category_id"`
}