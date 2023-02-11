package main

import (
	"database/sql"
	"net/http"

	"github.com/aslamaz/blood-donation/handler"
	"github.com/aslamaz/blood-donation/middleware"
	"github.com/aslamaz/blood-donation/repository"
	"github.com/go-chi/chi/v5"
	cm "github.com/go-chi/chi/v5/middleware"
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
	r.Use(cm.Logger) //logger is a global middleware so it will run for every end points below

	r.Get("/ping", handler.HandlePing)
	r.Post("/login", handler.HandleLogin)
	r.With(middleware.Authorize).Get("/user/me", handler.GetUser)
	r.Post("/user", handler.RegisterUser)
	r.With(middleware.Authorize).Patch("/user/me/password", handler.ChangePassword)
	r.With(middleware.Authorize).Get("/user/matching-blood-group", handler.GetMatchingBloodGroupsOfUser)
	r.Get("/blood-group/{bgid}/matching-blood-group", handler.GetMatchingBloodGroups)
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}

}
