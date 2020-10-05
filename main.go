package main

import (
	"fmt"
	"github.com/kons16/team7-backend/infra"
	"github.com/kons16/team7-backend/usecase"
	"github.com/kons16/team7-backend/web"
	"os"
)

func main() {
	dbMap, err := infra.NewDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	userRepo := infra.NewUserRepository(dbMap)
	userUC := usecase.NewUserUseCase(userRepo)

	s := web.NewServer(userUC)
	s.ListenAndServe()
}
