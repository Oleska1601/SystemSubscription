package entity

import "time"

// для норм отображения запроса get last subscription
type SubscriptionInfo struct {
	ID                   int64     `json:"subscription_id" example:"1"`
	StartDate            time.Time `json:"start_date" example:"2006-01-02T15:04:05Z"`
	EndDate              time.Time `json:"end_date" example:"2006-01-02T15:04:06Z"`
	SubscriptionTypeName string    `json:"subscription_type_name:" example:"1 second"`
	SubscriptionStatus   int       `json:"subscription_status" example:"1"` //1-active, 0-not active
}
