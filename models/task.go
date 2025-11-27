package models

import (
	"time"
)

type Task struct {
	Title       string
	Description string

	Status        bool
	Add_time      time.Time
	Compliet_time *time.Time
}

func NewTask(title, description string) *Task {
	return &Task{
		Title:       title,
		Description: description,

		Status:        false,
		Add_time:      time.Now(),
		Compliet_time: nil,
	}
}
