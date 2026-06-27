package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/whotterre/odysseus/src/internal/crypto"
	"github.com/whotterre/odysseus/src/internal/models"
	"github.com/whotterre/odysseus/src/internal/repositories"
)

type AuthService interface {
	AuthenticateUser(email, password string) (*models.User, error)
	RegisterUser(email, password, accountPubKey, encryptedAccountPrivKey, devicePubKey string) (string, error)
}

type authService struct {
	repo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) AuthService {
	return &authService{
		repo: repo,
	}
}

func (s *authService) AuthenticateUser(email, password string) (*models.User, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password cannot be empty")
	}

	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !crypto.CheckPasswordHash(password, user.PasswordHash) {
		return nil, errors.New("invalid email or password")
	}
	user.PasswordSalt = uuid.New().String()
	return user, nil
}

func (s *authService) RegisterUser(email, password, accountPubKey, encryptedAccountPrivKey, devicePubKey string) (string, error) {
	if email == "" || password == "" {
		return "", errors.New("email and password are required")
	}
	hashedPassword, err := crypto.HashPassword(password)
	if err != nil {
		return "", errors.New("failed to secure user credentials")
	}

	user := &models.User{
		Email:                      email,
		PasswordHash:               hashedPassword,
		AccountPublicKey:           accountPubKey,
		EncryptedAccountPrivateKey: encryptedAccountPrivKey,
		DevicePublicKey:            devicePubKey,
	}

	createdUser, err := s.repo.CreateUser(user)
	if err != nil {
		return "", err
	}

	return createdUser.ID.String(), nil
}
