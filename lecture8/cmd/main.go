package main

import (
	"lecture8/config"
	"lecture8/entity"
	"lecture8/handler"
	"lecture8/handler/httpserver"
	pg "lecture8/repository/pg"
	"lecture8/service"
	"log"
	"os"
	"os/signal"
)

func main() {

	cfg, err := config.Init("config.yaml")
	if err != nil {
		panic(err)
	}

	db, err := pg.New(
		pg.WithHost(cfg.DB.Host),
		pg.WithPort(cfg.DB.Port),
		pg.WithUsername(cfg.DB.Username),
		pg.WithPassword(cfg.DB.Password),
		pg.WithDBName(cfg.DB.DBName),
		pg.WithSSLMode(cfg.DB.SSLMode),
	)

	if err != nil {
		panic(err)
	}

	err = db.DB.AutoMigrate(&entity.User{}, &entity.Review{})
	if err != nil {
		panic(err)
	}

	defer db.Close()

	log.Println("db connection success")

	srvs := service.New(db)
	hndlr := handler.New(srvs)
	server := httpserver.New(
		hndlr.InitRouter(),
		httpserver.WithPort(cfg.HTTP.Port),
		httpserver.WithShutdownTimeout(cfg.HTTP.ShutdownTimeout),
	)

	log.Println("server started")
	server.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case s := <-interrupt:
		log.Printf("signal received: %s", s)
	case err = <-server.Notify():
		log.Printf("server notify: %s", err)
	}

	err = server.Shutdown()
	if err != nil {
		log.Printf("server shutdown err: %s", err)
	}

}
