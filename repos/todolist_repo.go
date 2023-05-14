package repos

import (
	"gorm.io/gorm"
)

// TodoRepo main repositories
type TodoRepo struct {
	DB *gorm.DB
}
