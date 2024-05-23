package services

import (
	"context"

	"github.com/chon26909/apartment-api/models"
	"github.com/chon26909/apartment-api/repositories"
)

type IUserService interface {
	GetProfile(ctx context.Context, userID int) (*models.ProfileResponse, error)
}

type UserService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(userRepository repositories.IUserRepository) IUserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) GetProfile(ctx context.Context, userId int) (*models.ProfileResponse, error) {

	profile, err := s.userRepository.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	return &models.ProfileResponse{
		Code:    1000,
		Message: "sucess",
		Data: models.Profile{
			Id:        profile.Id,
			FirstName: profile.FirstName,
			LastName:  profile.LastName,
			Email:     profile.Email,
		},
	}, nil
}
