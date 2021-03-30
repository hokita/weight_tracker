package domain

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Weight type
type Weight struct {
	ID        int       `json:"id"`
	Weight    int       `json:"weight"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Weights type
type Weights []Weight

// WeightRepository struct
type WeightRepository struct {
	DB *gorm.DB
}

// GetAll func
func (repo *WeightRepository) GetAll() (Weights, error) {
	var weights Weights

	result := repo.DB.Order("date desc").Find(&weights)
	if result.Error != nil {
		return Weights{}, result.Error
	}

	return weights, nil
}

// GetCurrents func
func (repo *WeightRepository) GetCurrents() Weights {
	today := repo.getTodayWeight()
	yesterday := repo.getYesterdayWeight()

	return []Weight{today, yesterday}
}

// Create func
func (repo *WeightRepository) Create(value int, date time.Time) error {
	weight := Weight{
		Weight: value,
		Date:   date,
	}

	result := repo.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "date"}},
		DoUpdates: clause.AssignmentColumns([]string{"weight"}),
	}).Create(&weight)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetTodayWeight func
func (repo *WeightRepository) getTodayWeight() Weight {
	var weight Weight
	repo.DB.First(&weight, "date = ?", time.Now())

	return weight
}

// GetYesterdayWeight func
func (repo *WeightRepository) getYesterdayWeight() Weight {
	var weight Weight
	repo.DB.First(&weight, "date = ?", time.Now().AddDate(0, 0, -1))

	return weight
}
