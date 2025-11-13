package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"-"` // hashed
	Projects  []Project `json:"projects" gorm:"many2many:user_projects;"`
	Tasks     []Task    `json:"tasks"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Project struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Tasks       []Task `json:"tasks"`
	Users       []User `json:"users" gorm:"many2many:user_projects;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Task struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	UserID      uint   `json:"user_id"`
	ProjectID   uint   `json:"project_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
