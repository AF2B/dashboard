package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ActivityVisitor struct {
    ID             primitive.ObjectID      `json:"-" bson:"_id,omitempty" gorm:"primaryKey"`
    CurrentValue   int                     `json:"currentValue"`
    PercentageChange float64               `json:"percentageChange"`
    Date           time.Time               `json:"createdAt"`
}

type VisitorPerMinute struct {
    ID             uint    `json:"-" bson:"_id,omitempty" gorm:"primaryKey"`
    CurrentValue   int
    PercentageChange float64
    Date           time.Time
}
