package routes

import (
	"go-backend/internal/handlers/admin"

	"github.com/gorilla/mux"
)

func AdminRoutes(r *mux.Router) {
	adminRouter := r.PathPrefix("/admin").Subrouter()
	//=========================Admin routes ==============
	adminRouter.HandleFunc("/questions", admin.GetAllQuestionsAdmin).Methods("GET")
}
