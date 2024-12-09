package service

import (
	"github.com/romulosm/go-modular-backend-template/internal/user/domain"
	"github.com/romulosm/go-modular-backend-template/internal/user/repository"
	"github.com/romulosm/go-modular-backend-template/pkg/logger"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *domain.User) error {
	logger.Log.Infof("Criando novo usuário: %s", user.Email)
	err := s.repo.Create(user)
	if err != nil {
		logger.Log.Errorf("Erro ao criar usuário: %v", err)
		return err
	}
	logger.Log.Infof("Usuário criado com sucesso: %s", user.ID)
	return nil
}

func (s *UserService) GetUser(id string) (*domain.User, error) {
	logger.Log.Infof("Buscando usuário com ID: %s", id)
	user, err := s.repo.GetByID(id)
	if err != nil {
		logger.Log.Errorf("Erro ao buscar usuário: %v", err)
		return nil, err
	}
	logger.Log.Infof("Usuário encontrado: %s", user.Email)
	return user, nil
}
