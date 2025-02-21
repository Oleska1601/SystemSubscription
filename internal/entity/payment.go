package entity

import "time"

type Payment struct {
	ID             int64
	Token          int64
	Status         string
	StartTime      time.Time
	EntTime        time.Time
	SubscriptionID int64
	UserID         int64
}
