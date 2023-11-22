package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"lecture10/handler"
	"lecture10/repo/cache"
	"lecture10/repo/pg"
	"log"
	"net/http"
)

func main() {
	db, err := pg.New(
		pg.WithHost("localhost"),
		pg.WithPort("5432"),
		pg.WithUsername("postgres"),
		pg.WithPassword("Alfarabi2004"),
		pg.WithDBName("postgres"),
		pg.WithSSLMode("disable"),
	)
	if err != nil {
		panic(err)
	}

	log.Println("db connection success")
	defer db.Close()

	redisClient, err := cache.NewRedisClient()
	if err != nil {
		panic(err)
	}
	bookCache := cache.NewBookCache(cache.BookCacheTimeout, redisClient)

	hndlr := handler.New(db, bookCache)

	r := gin.Default()
	r.GET("/books", hndlr.Get)
	r.POST("/books", hndlr.Create)

	go func() {
		r.Run(":8080")
	}()

	chiRouter := chi.NewRouter()
	chiRouter.Mount("/debug", middleware.Profiler())
	go func() {
		http.ListenAndServe(":8081", chiRouter)
	}()
}
