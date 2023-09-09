package service

import (
	"go-admin/modules/helper"
	"go-admin/modules/primitive/dto"
	models "go-admin/modules/primitive/model"
	"go-admin/modules/repository"
)

type AuthService struct {
	userRepository repository.UserRepository
}

func NewAuthService(repository repository.UserRepository) *AuthService {
	return &AuthService{
		userRepository: repository,
	}
}

func (s *AuthService) FindUserByEmail(email string) (*models.User, error) {
	data, err := s.userRepository.FindWhere(email)
	if err != nil {
		return nil, err
	} else {
	}
	return &data, nil
}

func (s *AuthService) FindUserBranch(id string) (*models.UserDetail, error) {
	data, err := s.userRepository.FindBranch(id)
	if err != nil {
		return nil, err
	} else {
	}
	return &data, nil
}

func (s *AuthService) RegisterUser(dto dto.UserDto) (*models.User, error) {
	hashPassword, err := helper.HashPassword(dto.Password)
	if err != nil {
		return nil, err
	}
	entity := models.User{
		Nik:      dto.Nik,
		Email:    dto.Email,
		Password: hashPassword,
		IsActive: dto.IsActive,
		TypeUser: dto.TypeUser,
	}
	data, err := s.userRepository.Save(entity)
	if err != nil {
		return nil, err
	} else {
		return &data, nil
	}
}
