package entity

import "time"

type Subscription struct {
	ID                 int64     `json:"subscription_id" example:"1"`
	Status             string    `json:"status" example:"active"` //active or not active
	StartDate          time.Time `json:"start_date" example:"2006-01-02T15:04:05Z"`
	EndDate            time.Time `json:"end_date" example:"2006-01-02T15:04:06Z"`
	SubscriptionTypeID int64     `json:"subscription_type_id" example:"1"`
	UserID             int64     `json:"user_id" example:"1"`
}
