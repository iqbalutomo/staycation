package service

import (
	"errors"
	"os"
	model "staycation/internal/models"
	repository "staycation/internal/repositories"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(req model.User) (*model.User, error)
	Login(req model.UserLoginRequest) (*model.LoginResponse, error)
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

func (s *authService) Login(req model.UserLoginRequest) (*model.LoginResponse, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return nil, err
	}

	resp := &model.LoginResponse{
		User:  user,
		Token: tokenString,
	}

	return resp, nil
}
