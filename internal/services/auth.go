package service

import (
	"errors"
	model "staycation/internal/models"
	repository "staycation/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(req model.User) (*model.User, error)
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo}
}

func (s *authService) Register(req model.User) (*model.User, error) {
	emailExist, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	} else if emailExist != nil {
		return nil, errors.New("email_exist")
	}

	phoneExist, err := s.repo.FindByPhone(req.Phone)
	if err != nil {
		return nil, err
	} else if phoneExist != nil {
		return nil, errors.New("phone_exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Phone:    req.Phone,
		Role:     model.UserRoleEnum(req.Role),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
