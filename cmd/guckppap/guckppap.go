package main

import (
	"context"
	"fmt"
	handler "gukppap-backend/internal/adapter/handler/http/fiber"
	repository "gukppap-backend/internal/adapter/repository/mysql"
	"gukppap-backend/internal/core/service"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func main() {

	// 의존성 결합
	ctx := context.Background()
	db, err := repository.NewDB(ctx)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	urlRepo := repository.NewURLRepository(db)
	urlService := service.NewURLService(urlRepo)
	urlHandler := handler.NewURLHandler(urlService)

	router := handler.NewRouter(urlHandler)
	if err := router.Run(9190); err != nil {
		fmt.Println("bug!, ", err)
		os.Exit(1)
	}
}
