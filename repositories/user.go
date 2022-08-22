package repositories

import "time"

type User struct {
	UserID    int       `db:"user_id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	Name      string    `db:"name"`
	LastName  string    `db:"last_name"`
	StartDate string    `db:"start_date"`
	ExpDate   time.Time `db:"exp_date"`
	Role      string    `db:"role"`
	Status    int       `db:"status"`
}

type UserRepository interface {
	Login(User) (*User, error)
	GetAll() ([]User, error)
	GetByID(int) (*User, error)
	Create(User) (*User, error)
	Update(int, User) error
	Delete(int) error
}
