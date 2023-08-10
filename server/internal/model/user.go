package model

// ========== ===== ==========
type UserRoles string
type UserPermission string

// ========== User Model ==========
type User struct {
	ID       int    `gorm:"autoIncrement; NOT NULL; primarykey"`
	Name     string `gorm:"size:255"`
	Email    string `gorm:"NOT NULL; UNIQUE_INDEX"`
	Password string `gorm:"NOT NULL"`
}

// ========== Input ==========
type CreateUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserPayload struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type PageInfo struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

type UsersConnection struct {
	Edges    []*User   `json:"edges"`
	PageInfo *PageInfo `json:"page_info"`
}

type AccessToken struct {
	Token string `json:"token"`
}

type UserCredentials struct {
	Email       string          `json:"email"`
	Permissions *UserPermission `json:"permissions"`
	Roles       *UserRoles      `json:"roles"`
}
