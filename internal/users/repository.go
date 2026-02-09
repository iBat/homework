package users

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	Dbpool *pgxpool.Pool
	Logger *slog.Logger
}

func NewUserRepository(dbpool *pgxpool.Pool, logger *slog.Logger) *UserRepository {
	return &UserRepository{
		Dbpool: dbpool,
		Logger: logger,
	}
}

func (r *UserRepository) CreateUser(form RegisterForm) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	query := `INSERT INTO users (name, email, password, created_at) VALUES ($1, $2, $3, NOW()) RETURNING id`
	var id int
	err = r.Dbpool.QueryRow(
		context.Background(),
		query,
		form.Name,
		form.Email,
		string(hashedPassword),
	).Scan(&id)
	if err != nil {
		r.Logger.Error("Failed to create user", slog.String("error", err.Error()))
		return "", err
	}
	r.Logger.Info("User created successfully", slog.String("email", form.Email))

	return form.Email, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*User, error) {
	query := `SELECT id, name, email, password, created_at FROM users WHERE email=$1`
	row := r.Dbpool.QueryRow(context.Background(), query, email)

	var user User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		r.Logger.Error("Failed to get user by email", slog.String("error", err.Error()))
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) ValidateUserPassword(email string, password string) (*User, error) {
	user, err := r.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		r.Logger.Warn("Invalid password attempt", slog.String("email", email))
		return nil, err
	}
	r.Logger.Info("User authenticated successfully", slog.String("email", email))
	return user, nil
}
