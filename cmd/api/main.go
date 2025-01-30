package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handler)

	allowedOrigins := []string{"https://stage-zero-o95z.onrender.com", "http://localhost:4000", "*"}

	srv := http.Server{
		Addr:         ":4000",
		Handler:      enableCORS(mux, allowedOrigins),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	slog.Info("starting server", "addr", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

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

func enableCORS(next http.Handler, allowedOrigins []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Origins")
		w.Header().Add("Vary", "Access-Control-Request-Method")

		origin := r.Header.Get("Origin")
		if origin != "" {
			for _, allowedOrigin := range allowedOrigins {
				if origin == allowedOrigin {
					w.Header().Set("Access-Control-Allow-Origin", origin)

					if r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != "" {
						w.Header().Set("Access-Control-Request-Method", "OPTIONS, PATCH, PUT, DELETE")
						w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
						w.WriteHeader(http.StatusOK)
						return
					}
					break
				}
			}
		}
		next.ServeHTTP(w, r)
	})
}
