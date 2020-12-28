package http

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Weight struct {
	ID        int       `json:"id"`
	Weight    int       `json:"weight"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type getWeightHandler struct {
	DB *gorm.DB
}

func (h *getWeightHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var todayWeight Weight
	var YesterdayWeight Weight
	h.DB.First(&todayWeight, "date = ?", time.Now())
	h.DB.First(&YesterdayWeight, "date = ?", time.Now().AddDate(0, 0, 1))

	m := map[string]string{
		"TodayWeight":     strconv.Itoa(todayWeight.Weight),
		"YesterdayWeight": strconv.Itoa(YesterdayWeight.Weight),
	}

	tpl := template.Must(template.ParseFiles("templates/index.html"))
	tpl.Execute(w, m)
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
	h.DB.First(&YesterdayWeight, "date = ?", time.Now().AddDate(0, 0, 1))

	tpl := template.Must(template.ParseFiles("templates/index.html"))
	m := map[string]string{
		"TodayWeight":     strconv.Itoa(weight.Weight),
		"YesterdayWeight": strconv.Itoa(YesterdayWeight.Weight),
	}
	tpl.Execute(w, m)
}
