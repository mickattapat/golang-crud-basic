package services

import "time"

// Get all
type UserReponse struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Role     string `json:"role"`
	ExpDate  string `json:"exp_date"`
}

// New account
type NewUserRequest struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Name            string `json:"name"`
	LastName        string `json:"last_name"`
	ExpDate         string `json:"exp_date"`
	Status          int    `json:"status"`
	Role            string `json:"role"`
}

type NewUserResponse struct {
	Username string    `json:"username"`
	Name     string    `json:"name"`
	LastName string    `json:"last_name"`
	ExpDate  time.Time `json:"exp_date"`
	Role     string    `json:"role"`
}

// Update account
type UpdateUserRequest struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Name            string `json:"name"`
	LastName        string `json:"last_name"`
	ExpDate         string `json:"exp_date"`
	Status          int    `json:"status"`
	Role            string `json:"role"`
}

// Login
type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserService interface {
	Login(UserLoginRequest) (string, error)
	GetUsers() ([]UserReponse, error)
	GetUser(int) (*UserReponse, error)
	CreateUser(NewUserRequest) (*NewUserResponse, error)
	UpdateUser(int, UpdateUserRequest) error
	DeleteUser(int) error
}
