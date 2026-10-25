package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		fmt.Println("Error loading .env file")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /db-connection", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(os.Getenv("DATABASE_URL"))); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	if err := http.ListenAndServe(os.Getenv("PORT"), mux); err != nil {
		fmt.Println(err)
	}
}
