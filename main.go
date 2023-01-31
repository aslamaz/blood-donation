package main

import (
	"database/sql"
	"net/http"

	"github.com/aslamaz/blood-donation/handler"
	"github.com/aslamaz/blood-donation/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:root@/blood_donation?parseTime=true")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	repository.Db = db

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/ping", handler.HandlePing)
	r.Post("/login", handler.HandleLogin)
	r.Get("/user/me", handler.GetUser)
	r.Post("/user", handler.RegisterUser)

	err = http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}

}
