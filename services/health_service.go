package services

import "github.com/chon26909/apartment-api/models"

type IHealthService interface {
	HealthCheck() (models.HealthCheckResponse, error)
}

type healthService struct {
}

func NewHealthService() IHealthService {
	return healthService{}
}

func (h healthService) HealthCheck() (models.HealthCheckResponse, error) {
	return models.HealthCheckResponse{Message: "success"}, nil
}
