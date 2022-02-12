package entities

import (
	"github.com/scyanh/velocitylimits/utils"
	"time"
)

type AttemptRequest struct {
	ID         int       `json:"id"`
	InputID    string    `json:"input_id"`
	CustomerID string    `json:"customer_id"`
	LoadAmount float64   `json:"load_amount"`
	Time       time.Time `json:"time"`
	Accepted   bool      `json:"accepted"`
}

type AttemptResponse struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	Accepted   bool   `json:"accepted"`
}

// Filter validates the attempt request
func (a *AttemptRequest) Filter() AttemptResponse {
	res := AttemptResponse{
		ID:         a.InputID,
		CustomerID: a.CustomerID,
		Accepted:   false,
	}

	count, validAttemptPerDay := a.ValidAttemptPerDay()

	// A maximum of 3 loads can be performed per day, regardless of amount
	if !validAttemptPerDay || count > utils.MaxAttemptsPerDay {
		return res
	}

	if !a.ValidAttemptPerWeek() {
		return res
	}

	res.Accepted = true
	a.Accepted = true

	return res
}

// ValidAttemptPerDay filters a maximum amount that can be loaded per week
func (a *AttemptRequest) ValidAttemptPerDay() (int64, bool) {
	inputs, count := NewLoadInputRepository().FindInputsPerDay(*a)
	totalAmount := 0.0
	for _, el := range inputs {
		totalAmount += el.LoadAmount
	}

	totalAmount += a.LoadAmount
	if !validAmountPerDay(totalAmount) {
		return count, false
	}

	return count, true
}

// ValidAttemptPerWeek filters a maximum amount that can be loaded per week
func (a *AttemptRequest) ValidAttemptPerWeek() bool {
	inputs := NewLoadInputRepository().FindInputsPerWeek(*a)
	totalAmount := 0.0
	for _, el := range inputs {
		totalAmount += el.LoadAmount
	}

	totalAmount += a.LoadAmount
	if !validAmountPerWeek(totalAmount) {
		return false
	}

	return true
}

// validAmountPerDay validates if the given amount is valid for the daily limit
func validAmountPerDay(amount float64) bool {
	if amount > utils.MaxAmountPerDay {
		return false
	}

	return true
}

// validAmountPerWeek validates if the given amount is valid for the weekly limit
func validAmountPerWeek(amount float64) bool {
	if amount > utils.MaxAmountPerWeek {
		return false
	}

	return true
}
