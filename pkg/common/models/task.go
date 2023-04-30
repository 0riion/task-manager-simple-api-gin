package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string    `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Details   string    `json:"details"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(fullName string, email string) *User {
	return &User{
		ID:        uuid.New().String(),
		FullName:  fullName,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewTask(title string, details string, startDate time.Time, endDate time.Time, user User) *Task {
	return &Task{
		ID:        uuid.New().String(),
		Title:     title,
		Details:   details,
		StartDate: startDate,
		EndDate:   endDate,
		User:      user,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
