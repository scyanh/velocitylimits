package entities

import (
	"github.com/kofoworola/godate"
	"github.com/scyanh/velocitylimits/db"
	"time"
)

type attemptRequestRepository struct{}

func NewLoadInputRepository() attemptRequestRepository {
	return attemptRequestRepository{}
}

// Save is used to save an attempt
func (a attemptRequestRepository) Save(inputRequest AttemptRequest) (*AttemptRequest, error) {
	if err := db.GetDb().Create(&inputRequest).Error; err != nil {
		return nil, err
	}
	return &inputRequest, nil
}

// FindAll finds all saved attempts
func (a attemptRequestRepository) FindAll() ([]AttemptRequest, int64) {
	var attempts []AttemptRequest

	var count int64
	db.GetDb().Find(&attempts).Count(&count)

	return attempts, count
}

// FindInput check if input attempt is repeated
func (a attemptRequestRepository) FindInput(req AttemptRequest) *AttemptRequest {
	var attempt AttemptRequest

	if db.GetDb().
		Where("input_id = ? AND customer_id = ?", req.InputID, req.CustomerID).
		First(&attempt).RowsAffected == 0 || attempt.ID == 0 {
		return nil
	}

	return &attempt
}

// FindInputsPerDay returns the success attempts made in a day for a specific customer
func (a attemptRequestRepository) FindInputsPerDay(req AttemptRequest) ([]AttemptRequest, int64) {
	var attempts []AttemptRequest

	today := godate.Create(req.Time)

	startDate := today.EndOfDay().Sub(1, godate.DAY).Format("2006-01-02T15:04:05Z")
	endDate := today.StartOfDay().Add(1, godate.DAY).Format("2006-01-02T15:04:05Z")

	var count int64
	db.GetDb().
		Where("accepted = true AND customer_id = ? AND time > ? AND time <= ?", req.CustomerID, startDate, endDate).
		Find(&attempts).
		Count(&count)

	return attempts, count
}

// FindInputsPerWeek returns the success attempts made in a week for a specific customer
func (a attemptRequestRepository) FindInputsPerWeek(req AttemptRequest) []AttemptRequest {
	var attempts []AttemptRequest

	today := godate.Create(req.Time)
	today.SetFirstDay(time.Monday)

	startDate := today.StartOfWeek().Format("2006-01-02T15:04:05Z")
	endDate := today.EndOfWeek().Format("2006-01-02T15:04:05Z")

	db.GetDb().
		Where("accepted = true AND customer_id = ? AND time BETWEEN ? AND ?", req.CustomerID, startDate, endDate).
		Find(&attempts)

	return attempts
}
