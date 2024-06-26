package postgres

import (
	"context"
	"fmt"

	"github.com/p3rfect05/grpc_service/internal/config"
	"github.com/p3rfect05/grpc_service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func New(dsn config.DatabaseURLConfig) (*Storage, error) {
	url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dsn.Host, dsn.User, dsn.Password, dsn.DBName, dsn.Port)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Storage{
		db: db,
	}, nil
}

func (s *Storage) SaveUser(ctx context.Context, email string, passHash []byte) (ID int64, err error) {
	user := models.User{Email: email, PasswordHash: passHash}
	const op = "storage.postgres.SaveUser"
	res := s.db.Create(&user)
	if res.Error != nil {
		return 0, fmt.Errorf("%s: %w", op, res.Error)
	}
	return user.ID, nil

}

func (s *Storage) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	const op = "storage.postgres.GetUserByEmail"
	res := s.db.First(&user)
	if res.Error != nil {
		return nil, fmt.Errorf("%s: %w", op, res.Error)
	}
	return &user, nil

}

func (s *Storage) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	return true, nil
}
