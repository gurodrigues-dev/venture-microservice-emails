package main

import (
	"log"

	"github.com/gurodrigues-dev/venture-microservice-emails/config"
	"github.com/gurodrigues-dev/venture-microservice-emails/internal/consumer"
	"github.com/gurodrigues-dev/venture-microservice-emails/internal/repository"
	"github.com/gurodrigues-dev/venture-microservice-emails/internal/service"
)

func main() {

	config, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatalf("error loading config: %s", err.Error())
	}

	repo, err := repository.NewPostgres()
	if err != nil {
		log.Fatalf("error creating repository: %s", err.Error())
	}

	aws, err := repository.NewAwsConnection()
	if err != nil {
		log.Fatalf("error creating aws connection: %s", err.Error())
	}

	service := service.New(repo, aws)

	consumer := consumer.New(service)

	log.Printf("initing service: %s", config.Name)
	consumer.Start()

}
