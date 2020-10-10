package usecase

import (
	"fmt"
	"github.com/kons16/team7-backend/domain/entity"
	"github.com/kons16/team7-backend/domain/repository"
	"github.com/kons16/team7-backend/domain/service"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	userRepo    repository.User
	sessionRepo repository.Session
}

// ユーザー登録する際に使用する構造体
type User struct {
	Name     string
	SubmitID string
	Year     int
	Sex      int
	Password string
}

func NewUserUseCase(userRepo repository.User, sessionRepo repository.Session) *UserUseCase {
	return &UserUseCase{userRepo: userRepo, sessionRepo: sessionRepo}
}

func (uc *UserUseCase) CreateUser(user *User) (string, error) {
	var createUserModel entity.User
	createUserModel.SubmitID = user.SubmitID
	createUserModel.Name = user.Name
	createUserModel.Year = user.Year
	createUserModel.Sex = user.Sex

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	createUserModel.PasswordHash = string(passwordHash)

	// MySQL にデータを保存
	userID, err := uc.userRepo.Create(&createUserModel)
	if err != nil {
		fmt.Print(err)
		return "", err
	}

	// SessionIDを生成
	sessionID := service.CreateNewToken()

	// Redis にセッションを保存
	err = uc.sessionRepo.CreateUserSession(userID, sessionID)
	if err != nil {
		fmt.Print(err)
		return "", err
	}

	return sessionID, nil
}
