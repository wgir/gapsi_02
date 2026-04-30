package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/user/gapsi_orders_api/internal/domain"
	"github.com/user/gapsi_orders_api/internal/infra/database/sqlc"
)

type userRepo struct {
	queries UserQuerier
}

func NewUserRepository(q UserQuerier) domain.UserRepository {
	return &userRepo{queries: q}
}

func (r *userRepo) Create(ctx context.Context, user *domain.User) error {
	role := sqlc.UserRoleUSER
	if user.Role == domain.RoleAdmin {
		role = sqlc.UserRoleADMIN
	}

	dbUser, err := r.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Role:         role,
	})
	if err != nil {
		return err
	}

	user.ID = dbUser.ID.String()
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
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	dbUser, err := r.queries.GetUserByID(ctx, uid)
	if err != nil {
		return nil, err
	}
	return mapDBUserToDomain(dbUser), nil
}

func mapDBUserToDomain(dbUser sqlc.User) *domain.User {
	role := domain.RoleUser
	if dbUser.Role == sqlc.UserRoleADMIN {
		role = domain.RoleAdmin
	}
	return &domain.User{
		ID:           dbUser.ID.String(),
		Email:        dbUser.Email,
		PasswordHash: dbUser.PasswordHash,
		Role:         role,
		CreatedAt:    dbUser.CreatedAt,
		UpdatedAt:    dbUser.UpdatedAt,
	}
}
