package admin
import (
	"encoding/json"
	"net/http"
	"go-backend/internal/db"
	"go-backend/internal/repositories"
)

func GetAllQuestionsAdmin(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","application/json")
	question, err  := repositories.GetAllQuestions(db.DB)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(question)
}