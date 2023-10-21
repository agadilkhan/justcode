package main

import (
	"github.com/gin-gonic/gin"
	"lecture10/handler"
	"lecture10/repo/cache"
	"lecture10/repo/pg"
	"log"
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
	r.GET("/", hndlr.Get)
	log.Println(r.Run(":8080"))
	//r.POST("/", hndlr.Create)
}
