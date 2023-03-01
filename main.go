package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pzentenoe/onboarding-go/app/shared/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	carUsecase "github.com/pzentenoe/onboarding-go/app/application/usecases/car"
	carMongoRepository "github.com/pzentenoe/onboarding-go/app/infrastructure/database/mongo/car"
	v1 "github.com/pzentenoe/onboarding-go/app/infrastructure/rest/v1"
	"github.com/pzentenoe/onboarding-go/app/infrastructure/rest/v1/car"
)

func main() {

	log.InitLogger()

	e := echo.New()
	e.Use(middleware.CORS())
	e.HideBanner = true
	g := e.Group("/v1")

	v1.NewHealth(g)

	mongoClient := createMongoClient()

	mongoRepository := carMongoRepository.NewCarMongoRepository(mongoClient)
	usecaseCar := carUsecase.NewCarUsecase(mongoRepository)
	car.NewCarHandler(usecaseCar, g)

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	go startServer(e, exit)
	<-exit
	gracefulShutdown(e)
}

func startServer(e *echo.Echo, exit chan os.Signal) {
	log.GetLogger().Infof("stating server at port :%d", 8080)
	err := e.Start(":8080")
	if err != nil {
		e.Logger.Fatal()
		close(exit)
	}

}

func gracefulShutdown(e *echo.Echo) {
	//TODO: Detener BD, Cerrar conexiones
	fmt.Println("Shutting down gracefully")
	e.Shutdown(context.Background())
}

func createMongoClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.GetLogger().Errorf(err.Error())
	}
	return client
}
