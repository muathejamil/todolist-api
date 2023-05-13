package repo

import (
	"gorm.io/gorm"
)

type TodoRepo struct {
	DB *gorm.DB
}
