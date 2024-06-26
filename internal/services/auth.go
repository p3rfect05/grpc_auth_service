package auth

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/p3rfect05/grpc_service/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	logger   *log.Logger
	storage  Storage
	tokenTTL time.Duration
}

type Storage interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (ID int64, err error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

func New(logger *log.Logger, storage Storage, tokenTTl time.Duration) *Auth {
	return &Auth{
		logger:   logger,
		storage:  storage,
		tokenTTL: tokenTTl,
	}
}

func (a *Auth) Login(ctx context.Context, email string, password string) (token string, err error) {
	user, err := a.storage.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password)); err != nil {
		return "", fmt.Errorf("invalid credentials")
	}
	return user.Email, nil
}

func (a *Auth) RegisterNewUser(ctx context.Context, email string, password string) (int64, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	userID, err := a.storage.SaveUser(ctx, email, hashPass)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (a *Auth) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	return true, nil
}
