package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/labstack/echo"
	"github.com/thiagoluiznunes/ze-challenge/data"
	"github.com/thiagoluiznunes/ze-challenge/infra/config"
	"github.com/thiagoluiznunes/ze-challenge/server"
	"github.com/thiagoluiznunes/ze-challenge/server/router"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		fmt.Println("couldn't read config file")
		return
	}
	_, err = json.Marshal(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("connecting to the database at %s:%s.", cfg.DBHost, cfg.DBPort)
	db, err := data.Connect(*(cfg))
	if err != nil {
		fmt.Println(err)
	}

	atInterruption(func() {
		log.Printf("closing database connection.")
		db.Close()
	})

	e := echo.New()
	srv := server.Instance(e, cfg)

	appRouter := router.New(srv.Echo, cfg, "ze-delivery")
	srv.AddAppRouter(appRouter)
}

func atInterruption(fn func()) {
	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, os.Interrupt)
		<-sc

		fn()
		os.Exit(0)
	}()
}
