package main

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	ID          uint   `gorm:"primary_key" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"done"`
}
