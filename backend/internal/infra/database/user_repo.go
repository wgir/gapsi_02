package database

import (
	"context"

	"github.com/user/gapsi_orders_api/internal/domain"
)

type userRepo struct {
	queries UserQuerier
}

func NewUserRepository(q UserQuerier) domain.UserRepository {
	return &userRepo{queries: q}
}

func (r *userRepo) Create(ctx context.Context, user *domain.User) error {
	role := UserRoleUSER
	if user.Role == domain.RoleAdmin {
		role = UserRoleADMIN
	}

	dbUser, err := r.queries.CreateUser(ctx, CreateUserParams{
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Role:         role,
	})
	if err != nil {
		return err
	}

	user.ID = dbUser.ID
	user.CreatedAt = dbUser.CreatedAt
	user.UpdatedAt = dbUser.UpdatedAt
	return nil
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	dbUser, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return mapDBUserToDomain(dbUser), nil
}

func (r *userRepo) GetByID(ctx context.Context, id string) (*domain.User, error) {
	dbUser, err := r.queries.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapDBUserToDomain(dbUser), nil
}

func mapDBUserToDomain(dbUser User) *domain.User {
	role := domain.RoleUser
	if dbUser.Role == UserRoleADMIN {
		role = domain.RoleAdmin
	}
	return &domain.User{
		ID:           dbUser.ID,
		Email:        dbUser.Email,
		PasswordHash: dbUser.PasswordHash,
		Role:         role,
		CreatedAt:    dbUser.CreatedAt,
		UpdatedAt:    dbUser.UpdatedAt,
	}
}
