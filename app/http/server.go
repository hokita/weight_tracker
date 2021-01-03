package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start() error {
	fmt.Println("start server")

	dns := "host=db user=app dbname=weight_tracker password=password sslmode=disable"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return err
	}

	r := mux.NewRouter()
	r.Handle("/", &getWeightHandler{db}).Methods("GET")
	r.Handle("/weights/", &createWeightHandler{db}).Methods("POST")
	r.Handle("/weights/all/", &getAllWeightsHandler{db}).Methods("GET")

	if err := http.ListenAndServe(":8080", r); err != nil {
		return err
	}

	return nil
}
