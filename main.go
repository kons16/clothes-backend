package main

import (
	"fmt"
	"github.com/kons16/team7-backend/infra"
	"github.com/kons16/team7-backend/usecase"
	"os"
)

func main() {
	dbMap, err := infra.NewDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	userRepo := infra.NewUserRepository(dbMap)
	_ = usecase.NewUserUseCase(userRepo)
}
