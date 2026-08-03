package service

import (
	"errors"
	"strings"

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

	// Validate the request
	// TODO: Move this validation to a separate function
	if req.Name == "" {
		return errors.New("name is required")
	}

	if req.Email == "" {
		return errors.New("email is required")
	}

	if strings.Contains(req.Email, "@") == false {
		return errors.New("invalid email")
	}

	if req.Password == "" {
		return errors.New("password is required")
	}

	if len(req.Password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	// Check if the email already exists
	_, err := s.repo.FindByEmail(req.Email)

	if err == nil {
		return errors.New("email already exists")
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	// Create the user
	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hash),
	}

	// Save the user to the database
	return s.repo.Create(user)
}
