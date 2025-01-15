package models

import "time"

type MemeCoin struct {
	ID              string    `json:"id" bson:"_id,omitempty"`
	Name            string    `json:"name" bson:"name" validate:"required,min=1"`
	Description     string    `json:"description,omitempty" bson:"description"`
	PopularityScore int       `json:"popularity_score" bson:"popularity_score"`
	CreatedAt       time.Time `json:"created_at" bson:"created_at"`
}
