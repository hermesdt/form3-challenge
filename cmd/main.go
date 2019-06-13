package main

import (
	"fmt"
	"net/http"

	app "github.com/hermesdt/form3-challenge"
)

func main() {
	app := app.NewApp()
	port := 3000
	fmt.Println("Listening on port", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), app.Router())
}
