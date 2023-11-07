package app

import (
	"fmt"
	"lecture12/orders/internal/config"
	"lecture12/orders/internal/entity"
	"lecture12/orders/internal/handler"
	"lecture12/orders/internal/repository"
	"lecture12/orders/internal/service"
	"lecture12/pkg/database/postgres"
	"lecture12/pkg/server/httpserver"
	"log"
	"os"
	"os/signal"
)

func Run() error {
	cfg, err := config.New("orders/internal/config")
	if err != nil {
		return fmt.Errorf("reading from config err: %v", err)
	}

	db, err := postgres.New(cfg.URL)
	if err != nil {
		return fmt.Errorf("database connection err: %v", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("database close err: %v", err)
		}
	}()

	err = db.AutoMigrate(&entity.Order{}, &entity.Customer{}, &entity.OrderProduct{}, &entity.Product{})
	if err != nil {
		return fmt.Errorf("AutoMigrate err: %v", err)
	}

	log.Println("database connection success")

	repo := repository.NewRepository(db)
	srvc := service.NewService(repo)
	hndlr := handler.New(srvc, cfg)
	server := httpserver.New(
		hndlr.InitRouter(),
		httpserver.WithPort(cfg.Port),
		httpserver.WithReadTimeOut(cfg.ReadTimeout),
		httpserver.WithWriteTimeout(cfg.WriteTimeout),
		httpserver.WithShutDownTimeout(cfg.ShutdownTimeout),
	)

	server.Start()

	log.Printf("server started")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case s := <-interrupt:
		log.Printf("signal received: %s", s.String())
	case err = <-server.Notify():
		log.Printf("server notify err: %s", err.Error())
	}

	err = server.ShutDown()
	if err != nil {
		log.Printf("server shutdown err: %v", err)
	}

	return nil
}
