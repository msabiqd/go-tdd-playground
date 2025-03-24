package service

import "go-tdd-playground/internal/user/repository"

type UserService struct {
	UserRepository repository.UserRepositoryInterface
}

func NewUserService(userRepository repository.UserRepositoryInterface) *UserService {
	return &UserService{UserRepository: userRepository}
}
