package users

import "time"

type RegisterForm struct {
	Name     string
	Email    string
	Password string
}

type User struct {
	Id        int       `db:"id"`
	Email     string    `db:"email"`
	Name      string    `db:"name"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}
