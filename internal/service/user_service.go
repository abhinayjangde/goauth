package service

import (
	"errors"

	"github.com/abhinayjangde/goauth/internal/model"
	"github.com/abhinayjangde/goauth/internal/respository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *respository.UserRepository
}

func NewUserService(repo *respository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Register(req *model.RegisterRequest) error {
	if req.Name == "" {
		return errors.New("name is required")
	}

	if req.Email == "" {
		return errors.New("email is required")
	}

	if req.Password == "" {
		return errors.New("password is required")
	}
	_, err := s.repo.FindByEmail(req.Email)

	if err == nil {
		return errors.New("email already exists")
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hash),
	}

	return s.repo.Create(user)
}
