package entity

type SubscriptionStatus struct {
	ID     int64 `json:"subscription_status_id" example:"1"`
	Status int   `json:"subscription_status" example:"1"` // 1-active, 0-not active
}
