package auth

import (
	"context"
	"log"
	"time"

	"github.com/p3rfect05/grpc_service/internal/models"
)

type Auth struct {
	logger   *log.Logger
	storage  Storage
	tokenTTL time.Duration
}

type Storage interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (ID int, err error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
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

}

func (a *Auth) RegisterNewUser(ctx context.Context, email string, password string) (userID int64, err error) {

}

func (a *Auth) IsAdmin(ctx context.Context, userID int64) (bool, error) {

}
