package entity

import "time"

type Subscription struct {
	ID                 int64
	Status             string
	StartDate          time.Time
	EndDate            time.Time
	SubscriptionTypeID int64
	UserID             int64
}
