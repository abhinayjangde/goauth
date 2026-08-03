package service

import (
	"errors"
	"time"

	"github.com/abhinayjangde/goauth/internal/config"
	"github.com/abhinayjangde/goauth/internal/model"
	"github.com/abhinayjangde/goauth/internal/respository"
	"github.com/abhinayjangde/goauth/internal/utils"
	"github.com/abhinayjangde/goauth/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo        *respository.UserRepository
	cfg         *config.Config
	refreshRepo *respository.RefreshTokenRepository
}

func NewUserService(repo *respository.UserRepository, cfg *config.Config, refreshRepo *respository.RefreshTokenRepository) *UserService {
	return &UserService{
		repo:        repo,
		cfg:         cfg,
		refreshRepo: refreshRepo,
	}
}

func (s *UserService) Register(req *model.RegisterRequest) error {

	err := validator.Validate.Struct(req)

	if err != nil {
		return err
	}

	// Check if the email already exists
	_, err = s.repo.FindByEmail(req.Email)

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

func (s *UserService) Login(req *model.LoginRequest) (*model.LoginResponse, error) {
	err := validator.Validate.Struct(req)
	if err != nil {
		return nil, err
	}
	user, err := s.repo.FindByEmail(req.Email)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	accessToken, err := utils.GenerateToken(
		user.ID,
		user.Email,
		s.cfg.JwtSecret,
	)

	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	newToken := &model.RefreshToken{
		UserId:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	if err := s.refreshRepo.Save(newToken); err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *UserService) Refresh(req *model.RefreshRequest) (*model.LoginResponse, error) {
	rt, err := s.refreshRepo.Find(req.RefreshToken)

	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	if time.Now().After(rt.ExpiresAt) {

		s.refreshRepo.Delete(req.RefreshToken)

		return nil,
			errors.New("refresh token expired")
	}

	err = s.refreshRepo.Delete(req.RefreshToken)

	if err != nil {
		return nil, err
	}

	user, err := s.repo.GetByID(rt.UserId)

	if err != nil {
		return nil, err
	}

	accessToken, err := utils.GenerateToken(
		user.ID,
		user.Email,
		s.cfg.JwtSecret,
	)

	refreshToken, err := utils.GenerateRefreshToken()

	newToken := &model.RefreshToken{
		UserId:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	err = s.refreshRepo.Save(newToken)

	return &model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
