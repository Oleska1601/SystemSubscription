package entity

import "time"

type Payment struct {
	ID                   int64     `json:"payment_id" example:"1"`
	Token                string    `json:"token"`
	SubscriptionTypeName string    `json:"subscription_type_name:" example:"1 second"`
	Amount               int64     `json:"amount" example:"10"`
	StartTime            time.Time `json:"start_time" example:"2006-01-02T15:04:05Z"`
	EntTime              time.Time `json:"end_time" example:"2006-01-02T15:04:15Z"`
	PaymentStatusID      int64     `json:"payment_status_id" example:"1"`
	UserID               int64     `json:"user_id" example:"1"`
}
