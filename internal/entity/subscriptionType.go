package entity

type SubscriptionType struct {
	ID       int64  `json:"subscription_type_id" example:"1"`
	TypeName string `json:"type_name" example:"1 second"`
	Duration int    `json:"duration" example:"1"` //1/3/6/12
	Price    int64  `json:"price" example:"10"`
}
