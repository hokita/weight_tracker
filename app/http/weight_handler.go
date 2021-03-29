package http

import (
	"encoding/json"
	"net/http"

	"github.com/hokita/weight_tracker/domain"

	"gorm.io/gorm"
)

type getWeightHandler struct {
	DB *gorm.DB
}

func (h *getWeightHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	repo := domain.WeightRepository{DB: h.DB}
	weights := repo.GetCurrents()

	if err := json.NewEncoder(w).Encode(weights); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type createWeightHandler struct {
	DB *gorm.DB
}

func (h *createWeightHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	var params struct {
		Weight int `json:"weight"`
	}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	repo := domain.WeightRepository{DB: h.DB}
	if err := repo.Create(params.Weight); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

type getAllWeightsHandler struct {
	DB *gorm.DB
}

func (h *getAllWeightsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	repo := domain.WeightRepository{DB: h.DB}
	weights, err := repo.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(weights); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
