package main

import (
	"lecture9/internal/entity"
	"lecture9/internal/handler"
	"lecture9/internal/repository/pg"
	"lecture9/internal/service"
	"lecture9/pkg/httpserver"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	db, err := pg.NewPostgres(
		pg.WithHost("localhost"),
		pg.WithPort("5432"),
		pg.WithUsername("postgres"),
		pg.WithPassword("Alfarabi2004"),
		pg.WithSSLMode("disable"),
	)
	if err != nil {
		panic(err)
	}

	err = db.DB.AutoMigrate(&entity.User{})
	if err != nil {
		panic(err)
	}

	defer db.Close()

	log.Println("db connection success")

	srvs := service.NewManager(db)

	hndlr := handler.New(srvs)

	server := httpserver.New(
		hndlr.Init(),
		httpserver.WithPort(":8080"),
		httpserver.WithShutdownTimeout(time.Duration(30)*time.Second),
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
