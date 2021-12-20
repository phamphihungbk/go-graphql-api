package model

// ========== User Model ==========
type User struct {
	ID       int    `gorm:"primarykey"`
	Name     string `gorm:"size:255"`
	Email    string `gorm:"NOT NULL; UNIQUE_INDEX"`
	Password string `gorm:"NOT NULL"`
}

// ========== Input ==========
type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
