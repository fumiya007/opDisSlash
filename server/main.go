package main

import (
	"net/http"
	"os"

	"github.com/fumiya007/tpDisSlash/api"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	http.HandleFunc("/", api.Handler)
	// ポートで起動
	http.ListenAndServe(":"+port, nil)
}
