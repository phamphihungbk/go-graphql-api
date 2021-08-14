package migrations

import "gorm.io/gorm"

type Test struct {
	gorm.Model
	Test  string `gorm:"size:255"`
}
