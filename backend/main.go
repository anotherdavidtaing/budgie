package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/anotherdavidtaing/budgie/internal/auth"
	"github.com/anotherdavidtaing/budgie/internal/database"
	"github.com/anotherdavidtaing/budgie/internal/env"
)

type Category struct {
	ID     int    `json:"category_id"`
	Name   string `json:"name"`
	UserId string `json:"user_id"`
}

func GetCategory(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM category")
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		defer rows.Close()

		var categories []Category

		for rows.Next() {
			var category Category
			err := rows.Scan(&category.ID, &category.Name, &category.UserId)
			if err != nil {
				log.Println(err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			categories = append(categories, category)
		}

		categoriesJSON, err := json.Marshal(categories)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Set the appropriate headers and write the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(categoriesJSON)
	}
}

func CreateCategory(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var category Category
		err := json.NewDecoder(r.Body).Decode(&category)
		if err != nil {
			log.Println(err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		log.Printf("Received user: %+v\n", category)

		stmt, err := db.Prepare("INSERT INTO category (name, user_id) VALUES ($1, $2)")
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		defer stmt.Close()

		_, err = stmt.Exec(category.Name, category.UserId)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func main() {
	env.LoadEnv()
	db := database.Connect()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	_, authMux, err := auth.New()
	if err != nil {
		log.Fatal(err)
	}

	r.Mount("/auth", authMux)

	r.Get("/", GetCategory(db))
	r.Post("/", CreateCategory(db))

	fmt.Println("Starting http server on PORT :3000")

	err = http.ListenAndServe(":3000", r)

	if err != nil {
		log.Fatalln("Failed to start server on PORT :3000")
	}
}
