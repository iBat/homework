package api

type RegisterForm struct {
	Name     string
	Email    string
	Password string
}

type LoginForm struct {
	Email    string
	Password string
}
