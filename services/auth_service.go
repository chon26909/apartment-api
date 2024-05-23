package services

import (
	"context"
	"fmt"
	"time"

	"github.com/chon26909/apartment-api/models"
	"github.com/chon26909/apartment-api/repositories"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type IAuthService interface {
	Login(ctx context.Context, request *models.LoginRequest) (*models.LoginResponse, error)
}

type AuthService struct {
	userRepository repositories.IUserRepository
}

func NewAuthService(userRepository repositories.IUserRepository) IAuthService {
	return &AuthService{userRepository: userRepository}
}

// Login implements IAuthService.
func (s *AuthService) Login(ctx context.Context, request *models.LoginRequest) (*models.LoginResponse, error) {

	user, err := s.userRepository.GetUserByEmail(request.Email)
	if err != nil {
		return nil, err
	}

	secertKey := []byte("abcd")

	now := time.Now()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, models.CustomClaims{
		UserId: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: "access-token",
			ExpiresAt: &jwt.NumericDate{
				Time: now.Add(time.Duration(5 * time.Minute)),
			},
			NotBefore: &jwt.NumericDate{
				Time: now,
			},
			IssuedAt: &jwt.NumericDate{
				Time: now,
			},
		},
	})

	tokenSigned, err := accessToken.SignedString(secertKey)
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, models.CustomClaims{
		UserId: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: "access-token",
			ExpiresAt: &jwt.NumericDate{
				Time: now.Add(time.Duration(10 * time.Minute)),
			},
			NotBefore: &jwt.NumericDate{
				Time: now,
			},
			IssuedAt: &jwt.NumericDate{
				Time: now,
			},
		},
	})

	refreshTokenSigned, err := refreshToken.SignedString(secertKey)
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		Code:    1000,
		Message: "success",
		Data: models.LoginResponseData{
			AccessToken:  tokenSigned,
			RefreshToken: refreshTokenSigned,
		},
	}, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, request *models.RequestTokenRequest) (*models.LoginResponse, error) {

	user, err := s.userRepository.GetUserById(1)
	if err != nil {
		return nil, err
	}

	userId := fmt.Sprintf("%d", user.Id)

	type CustomClaims struct {
		jwt.RegisteredClaims
	}

	// validate refresh token expire
	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(request.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return nil, echo.ErrUnauthorized

	})

	if err != nil || !token.Valid {
		return nil, echo.ErrUnauthorized
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, echo.ErrUnauthorized
	}

	secertKey := []byte("abcd")

	now := time.Now()

	// generate new access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject: "refresh-token",
		ExpiresAt: &jwt.NumericDate{
			Time: now.Add(time.Duration(30 * time.Minute)),
		},
		NotBefore: &jwt.NumericDate{
			Time: now,
		},
		IssuedAt: &jwt.NumericDate{
			Time: now,
		},
		ID: userId,
	})

	tokenSigned, err := accessToken.SignedString(secertKey)
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		Code:    1000,
		Message: "success",
		Data: models.LoginResponseData{
			AccessToken: tokenSigned,
		},
	}, nil
}
