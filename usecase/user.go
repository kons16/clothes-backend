package usecase

import "github.com/kons16/team7-backend/domain/repository"

type UserUseCase struct {
	userRepo repository.User
}

func NewUserUseCase(userRepo repository.User) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}
