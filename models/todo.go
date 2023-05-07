package models

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	gorm.Model
	Title       string     `json:"Title"`
	Description string     `json:"Description"`
	StartDay    *time.Time `json:"StartDay"`
	DueDay      *time.Time `json:"DueDay"`
}
