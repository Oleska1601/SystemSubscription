package entity

import "time"

type SubscriptionType struct {
	ID       int64
	TypeName string
	Duration time.Time
	Price    int64
}
