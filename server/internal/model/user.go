package model

type User struct {
	ID       int    `gorm:"primarykey"`
	Name     string `gorm:"size:255"`
	Email    string `gorm:"NOT NULL; UNIQUE_INDEX"`
	Password string `gorm:"NOT NULL"`
}
