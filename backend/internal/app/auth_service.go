package app

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/user/gapsi_orders_api/internal/domain"
	"github.com/user/gapsi_orders_api/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo        domain.UserRepository
	tokenHelper *jwt.TokenHelper
	accessTTL   time.Duration
	refreshTTL  time.Duration
}

func NewAuthService(repo domain.UserRepository, th *jwt.TokenHelper, accessTTL, refreshTTL time.Duration) domain.UserService {
	return &AuthService{
		repo:        repo,
		tokenHelper: th,
		accessTTL:   accessTTL,
		refreshTTL:  refreshTTL,
	}
}

func (s *AuthService) Register(ctx context.Context, email, password string, role domain.UserRole) (*domain.User, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Email:        email,
		PasswordHash: string(hashed),
		Role:         role,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "unique constraint") {
			return nil, errors.New("el correo electrónico ya está registrado")
		}
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, string, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return "", "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", "", errors.New("invalid credentials")
	}

	accessToken, err := s.tokenHelper.GenerateToken(user.ID, user.Email, string(user.Role), "access", s.accessTTL)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.tokenHelper.GenerateToken(user.ID, user.Email, string(user.Role), "refresh", s.refreshTTL)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, refreshToken string) (string, string, error) {
	claims, err := s.tokenHelper.ValidateToken(refreshToken)
	if err != nil || claims.Type != "refresh" {
		return "", "", errors.New("invalid refresh token")
	}

	// Double check user still exists
	user, err := s.repo.GetByID(ctx, claims.UserID)
	if err != nil {
		return "", "", errors.New("user not found")
	}

	accessToken, err := s.tokenHelper.GenerateToken(user.ID, user.Email, string(user.Role), "access", s.accessTTL)
	if err != nil {
		return "", "", err
	}

	newRefreshToken, err := s.tokenHelper.GenerateToken(user.ID, user.Email, string(user.Role), "refresh", s.refreshTTL)
	if err != nil {
		return "", "", err
	}

	return accessToken, newRefreshToken, nil
}

func (s *AuthService) GetMe(ctx context.Context, userID string) (*domain.User, error) {
	return s.repo.GetByID(ctx, userID)
}
