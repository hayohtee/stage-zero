package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var response struct {
		Email       string    `json:"email"`
		CurrentTime time.Time `json:"current_datetime"`
		GitHubURL   string    `json:"github_url"`
	}

	response.Email = "olamilekanakintilebo@gmail.com"
	response.CurrentTime = time.Now()
	response.GitHubURL = "https://github.com/hayohtee/stage-zero"

	js, err := json.MarshalIndent(&response, "", "\t")
	if err != nil {
		http.Error(w, "the server encountered a problem and could not process the request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
