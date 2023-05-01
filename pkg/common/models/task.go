package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FullName  string             `json:"full_name"`
	Email     string             `json:"email"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type Task struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title     string             `json:"title"`
	Details   string             `json:"details"`
	StartDate time.Time          `json:"start_date"`
	EndDate   time.Time          `json:"end_date"`
	User      User               `json:"user"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

func NewUser(fullName string, email string) *User {
	return &User{
		ID:        primitive.NewObjectID(),
		FullName:  fullName,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewTask(title string, details string, startDate time.Time, endDate time.Time, user User) *Task {
	return &Task{
		ID:        primitive.NewObjectID(),
		Title:     title,
		Details:   details,
		StartDate: startDate,
		EndDate:   endDate,
		User:      user,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
