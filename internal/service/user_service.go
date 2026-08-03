package service

import (
	"errors"
	"strings"

	"github.com/abhinayjangde/goauth/internal/config"
	"github.com/abhinayjangde/goauth/internal/model"
	"github.com/abhinayjangde/goauth/internal/respository"
	"github.com/abhinayjangde/goauth/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *respository.UserRepository
	cfg  *config.Config
}

func NewUserService(repo *respository.UserRepository, cfg *config.Config) *UserService {
	return &UserService{
		repo: repo,
		cfg:  cfg,
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

func (s *UserService) Login(req *model.LoginRequest) (string, error) {
	if req.Email == "" {
		return "", errors.New("email is required")
	}

	if req.Password == "" {
		return "", errors.New("password is required")
	}
	user, err := s.repo.FindByEmail(req.Email)

	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(
		user.ID,
		user.Email,
		s.cfg.JwtSecret,
	)

	if err != nil {
		return "", err
	}

	return token, nil
}
