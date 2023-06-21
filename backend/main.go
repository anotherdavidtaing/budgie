package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	fmt.Println("Starting http server on PORT :3000")

	err := http.ListenAndServe(":3000", r)

	if err != nil {
		log.Fatalln("Failed to start server on PORT :3000")
	}
}
