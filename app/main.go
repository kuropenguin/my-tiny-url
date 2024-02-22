package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/kuropenguin/my-tiny-url/app/handler"
	"github.com/kuropenguin/my-tiny-url/app/mysql"
	"github.com/kuropenguin/my-tiny-url/app/redis"
	"github.com/kuropenguin/my-tiny-url/app/repository"
	"github.com/kuropenguin/my-tiny-url/app/usecase"
)

func main() {
	// TODO
	// sqlc : mysql で sqlc 使う
	// request validation
	// logging

	log.Println("start server ...")

	router := mux.NewRouter()
	repo := repository.NewMysqlRepository(mysql.GetDB())
	cache := repository.NewCacheRedisRepository(redis.NewRedisClient())
	usecase := usecase.NewUsecaseImpl(repo, cache)
	handler := handler.NewHandlerImple(usecase)

	router.HandleFunc("/health", handler.Health).Methods("GET")
	router.HandleFunc("/create_tiny_url", handler.CreateTinyURL).Methods("POST")
	router.HandleFunc("/get_original_url", handler.GetOriginalURLByTinyURL).Methods("GET")

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
