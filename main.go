package main

import (
	"fmt"
	"github.com/kons16/team7-backend/infra/MySQL"
	"github.com/kons16/team7-backend/infra/Redis"
	"github.com/kons16/team7-backend/usecase"
	"github.com/kons16/team7-backend/web"
	"os"
)

func main() {
	dbMap, err := MySQL.NewMySQLDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	rdMap, err := Redis.NewRedisDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	userRepo := MySQL.NewUserRepository(dbMap)
	clothRepo := MySQL.NewClothRepository(dbMap)
	sessionRepo := Redis.NewSessionRepository(rdMap)

	userUC := usecase.NewUserUseCase(userRepo, sessionRepo)
	clothUC := usecase.NewClothUseCase(clothRepo, sessionRepo)
	sessionUC := usecase.NewSessionUseCase(sessionRepo)

	s := web.NewServer(userUC, sessionUC, clothUC)

	fmt.Println("Server Start!!")
	s.ListenAndServe()
}
