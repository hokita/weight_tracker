package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Start function
func Start() error {
	fmt.Println("start server")

	dns := "host=db user=app dbname=weight_tracker password=password sslmode=disable"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return err
	}

	r := mux.NewRouter()
	r.Handle("/", &getWeightHandler{db}).Methods(http.MethodGet)
	r.Handle("/", &createWeightHandler{db}).Methods(http.MethodPost, http.MethodOptions)
	r.Handle("/weights/all/", &getAllWeightsHandler{db}).Methods(http.MethodGet)
	r.Use(mux.CORSMethodMiddleware(r))

	if err := http.ListenAndServe(":8081", r); err != nil {
		return err
	}

	return nil
}
