package entities

import (
	"strconv"
	"time"
)

type Input struct {
	ID         string `json:"id"`
	CustomerId string `json:"customer_id"`
	LoadAmount string `json:"load_amount"`
	Time       string `json:"time"`
}

// Format is used to format date and amount and returns an attempt request
func (in Input) Format() (*AttemptRequest, error) {
	layout := "2006-01-02T15:04:05Z"
	t, err := time.Parse(layout, in.Time)
	if err != nil {
		return nil, err
	}

	amount, err := strconv.ParseFloat(in.LoadAmount[1:], 64)
	if err != nil {
		return nil, err
	}

	attemptRequest := AttemptRequest{
		InputID:    in.ID,
		CustomerID: in.CustomerId,
		LoadAmount: amount,
		Time:       t,
	}

	return &attemptRequest, nil
}
