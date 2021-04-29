package model

type User struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Password string `json:"password,omitempty"`
}

type UserInput struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Role  string `json:"role" example:"admin"`
}

type UserLogin struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

