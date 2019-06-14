package main

import (
	"fmt"
	"net/http"
	"os"

	app "github.com/hermesdt/form3-challenge"
)

func main() {
	app := app.NewApp()

	port := getPort()
	fmt.Println("Listening on port", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), app.Router())
}

func getPort() string {
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "3000"
	}

	return port
}
