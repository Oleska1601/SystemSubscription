package entity

type PaymentStatus struct {
	ID     int64  `json:"payment_status_id" example:"1"`
	Status string `json:"payment_status" example:"in process"` //success, on process, cancel
}
