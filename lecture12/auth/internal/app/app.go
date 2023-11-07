package app

import (
	"lecture12/auth/internal/config"
	"lecture12/auth/internal/entity"
	"lecture12/auth/internal/handler/http"
	"lecture12/auth/internal/repository"
	"lecture12/auth/internal/service"
	"lecture12/pkg/database/postgres"
	"lecture12/pkg/server/httpserver"
	"log"
	"os"
	"os/signal"
)

func Run() error {
	cfg, err := config.New("auth/internal/config")
	if err != nil {
		log.Panicf("reading from config err: %v", err)
	}

	db, err := postgres.New(cfg.URL)
	if err != nil {
		log.Panicf("database connection err: %v", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("database close err: %v", err)
		}
	}()

	err = db.AutoMigrate(&entity.User{}, &entity.Token{})
	if err != nil {
		log.Panicf("AutoMigrate err: %v", err)
	}

	log.Println("db connection success")

	repo := repository.NewRepository(db)
	srvc := service.NewService(repo, cfg)
	hndlr := http.New(srvc)
	server := httpserver.New(
		hndlr.InitRouter(),
		httpserver.WithPort(cfg.Port),
		httpserver.WithShutDownTimeout(cfg.ShutdownTimeout),
		httpserver.WithReadTimeOut(cfg.ReadTimeout),
		httpserver.WithWriteTimeout(cfg.WriteTimeout),
	)

	server.Start()

	log.Println("server started")

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
