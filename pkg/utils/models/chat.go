package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Users           []int64            `json:"users" bson:"users"`
	LastMessage     string             `json:"last_message" bson:"last_message"`
	LastMessageTime time.Time          `json:"last_message_time" bson:"last_message_time"`
}

type Message struct {
	Message string `json:"message"`
}
type UserShortDetail struct {
	Id             int64  `json:"id"`
	Username       string `json:"username"`
	ProfilePicture string `json:"profile_picture"`
}
