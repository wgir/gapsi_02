package domain

import (
	"context"
	"time"
)

type UserRole string

const (
	RoleAdmin UserRole = "ADMIN"
	RoleUser  UserRole = "USER"
)

type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Role         UserRole  `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByID(ctx context.Context, id string) (*User, error)
}

type UserService interface {
	Register(ctx context.Context, email, password string, role UserRole) (*User, error)
	Login(ctx context.Context, email, password string) (string, string, error) // accessToken, refreshToken, error
	RefreshToken(ctx context.Context, refreshToken string) (string, string, error)
	GetMe(ctx context.Context, userID string) (*User, error)
}
