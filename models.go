package main

import "time"

// Task model represents a task record.
type Task struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title" gorm:"size:255;not null"`
	Description string    `json:"description,omitempty" gorm:"type:text"`
	Status      string    `json:"status" gorm:"size:50;default:'todo'"` // todo | in_progress | done
	Priority    int       `json:"priority" gorm:"default:0"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
