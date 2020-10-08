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

	_, err = Redis.NewRedisDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	userRepo := MySQL.NewUserRepository(dbMap)
	userUC := usecase.NewUserUseCase(userRepo)

	s := web.NewServer(userUC)
	fmt.Println("Server Start!!")
	s.ListenAndServe()
}
