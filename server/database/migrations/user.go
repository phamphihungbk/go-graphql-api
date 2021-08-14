package migrations

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `gorm:"size:255"`
	Email     string `gorm:"NOT NULL; UNIQUE_INDEX"`
	Password  string `gorm:"NOT NULL"`
	Role      string `gorm:"NOT_NULL;size:255;DEFAULT:'standard'"`
	Active    bool   `gorm:"NOT NULL; DEFAULT: true"`
}
