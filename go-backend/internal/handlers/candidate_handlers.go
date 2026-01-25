package handlers

import (
	"encoding/json"
	"net/http"

)

func CandidateLogin(w http.ResponseWriter, r *http.Request){
     json.NewEncoder(w).Encode(map[string]string{
		"message":"Admin login API working",
	 })
}
