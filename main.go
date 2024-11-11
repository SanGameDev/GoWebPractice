package main

import (
	"log"
	"net/http"
	"time"

	"github.com/SanGameDev/GoWebPractice/internal/course"
	"github.com/SanGameDev/GoWebPractice/internal/enrollment"
	"github.com/SanGameDev/GoWebPractice/internal/user"

	"github.com/SanGameDev/GoWebPractice/pkg/bootstrap"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	router := mux.NewRouter()
	_ = godotenv.Load()
	l := bootstrap.InitLogger()

	db, err := bootstrap.DBConnection()
	if err != nil {
		l.Fatal(err)
	}

	userRepo := user.NewRepo(l, db)
	userSrv := user.NewService(l, userRepo)
	userEnd := user.MakeEndpoints(userSrv)

	courseRepo := course.NewRepo(db, l)
	courseSrv := course.NewService(l, courseRepo)
	courseEnd := course.MakeEndpoints(courseSrv)

	enrollRepo := enrollment.NewRepo(db, l)
	enrollSrv := enrollment.NewService(l, userSrv, courseSrv, enrollRepo)
	enrollEnd := enrollment.MakeEndpoints(enrollSrv)

	router.HandleFunc("/users", userEnd.Create).Methods("POST")
	router.HandleFunc("/users/{id}", userEnd.Get).Methods("GET")
	router.HandleFunc("/users", userEnd.GetAll).Methods("GET")
	router.HandleFunc("/users/{id}", userEnd.Update).Methods("PATCH")
	router.HandleFunc("/users/{id}", userEnd.Delete).Methods("DELETE")

	router.HandleFunc("/courses", courseEnd.Create).Methods("POST")
	router.HandleFunc("/courses/{id}", courseEnd.Get).Methods("GET")
	router.HandleFunc("/courses", courseEnd.GetAll).Methods("GET")
	router.HandleFunc("/courses/{id}", courseEnd.Update).Methods("PATCH")
	router.HandleFunc("/courses/{id}", courseEnd.Delete).Methods("DELETE")

	router.HandleFunc("/enrollments", enrollEnd.Create).Methods("POST")

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
