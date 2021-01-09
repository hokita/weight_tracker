package http

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Weight struct
type Weight struct {
	ID        int       `json:"id"`
	Weight    int       `json:"weight"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TopParams struct
type TopParams struct {
	Weight          Weight `json:"weight"`
	YesterdayWeight Weight `json:"yesterday_weight"`
}

type getWeightHandler struct {
	DB *gorm.DB
}

func (h *getWeightHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var todayWeight Weight
	var YesterdayWeight Weight
	h.DB.First(&todayWeight, "date = ?", time.Now())
	h.DB.First(&YesterdayWeight, "date = ?", time.Now().AddDate(0, 0, -1))

	tp := TopParams{
		Weight:          todayWeight,
		YesterdayWeight: YesterdayWeight,
	}
	if err := json.NewEncoder(w).Encode(tp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type createWeightHandler struct {
	DB *gorm.DB
}

func (h *createWeightHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	i, err := strconv.Atoi(r.FormValue("weight"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	weight := Weight{
		Weight: i,
		Date:   time.Now(),
	}

	result := h.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "date"}},
		DoUpdates: clause.AssignmentColumns([]string{"weight"}),
	}).Create(&weight)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var YesterdayWeight Weight
	h.DB.First(&YesterdayWeight, "date = ?", time.Now().AddDate(0, 0, -1))

	tpl := template.Must(template.ParseFiles("templates/index.html"))
	m := map[string]string{
		"Now":             time.Now().Format("2006-01-02 Mon"),
		"TodayWeight":     strconv.Itoa(weight.Weight),
		"YesterdayWeight": strconv.Itoa(YesterdayWeight.Weight),
	}
	tpl.Execute(w, m)
}

type getAllWeightsHandler struct {
	DB *gorm.DB
}

func (h *getAllWeightsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var weights []Weight
	result := h.DB.Order("date desc").Find(&weights)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tpl := template.Must(template.ParseFiles("templates/all.html"))
	m := map[string]interface{}{
		"Len":     len(weights),
		"Weights": weights,
	}
	tpl.Execute(w, m)
}
