package db

import (
	"github.com/jinzhu/gorm"
)

type BaseModel struct {
	gorm.Model
}
