package routes
import ( "github.com/gorilla/mux"
         "go-backend/internal/handlers"
)


func SetupRoutes() *mux.Router{

        r := mux.NewRouter()
        //=========================Admin routes ==============
        r.HandleFunc("/admin/login",handlers.AdminLogin).Methods("POST")


        // ================= CANDIDATE ROUTES =================
        r.HandleFunc("/candidate/login", handlers.CandidateLogin).Methods("POST")
	// r.HandleFunc("/candidate/start-exam", handlers.StartExam).Methods("POST")
	// r.HandleFunc("/candidate/submit", handlers.SubmitAnswers).Methods("POST")
	// r.HandleFunc("/candidate/result", handlers.GetResult).Methods("GET")
        return  r
}