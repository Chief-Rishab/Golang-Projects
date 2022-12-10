package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Article struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty"`
	Description string             `json:"description,omitempty"`
	URL         string             `json:"url,omitempty"`
	ImageURL    string             `json:"image,omitempty"`
	Read        bool               `json:"read,omitempty"`
}

type Message struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
	Image       string `json:"image,omitempty"`
}


