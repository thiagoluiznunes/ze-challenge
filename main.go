package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/thiagoluiznunes/ze-challenge/data"
	"github.com/thiagoluiznunes/ze-challenge/infra/config"
	"github.com/thiagoluiznunes/ze-challenge/server"
	"github.com/thiagoluiznunes/ze-challenge/server/router"
	"github.com/thiagoluiznunes/ze-challenge/server/router/partnerroute"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Error("couldn't read config file")
		return
	}
	_, err = json.Marshal(cfg)
	if err != nil {
		log.Error(err)
		return
	}

	log.Warn(fmt.Sprintf("connecting to the database at %s:%s.", cfg.DBHost, cfg.DBPort))
	db, err := data.Connect(*(cfg))
	if err != nil {
		fmt.Println(err)
	}

	atInterruption(func() {
		log.Printf("closing database connection.")
		db.Close()
	})

	e := echo.New()

	// Add controllers
	partnerController := partnerroute.NewController(cfg)

	// Initialize Routers
	partnerRoute := partnerroute.NewRouter("partner", partnerController)

	srv := server.Instance(e, cfg)

	appRouter := router.New(srv.Echo, cfg, "ze-delivery")
	appRouter.AddRouters(partnerRoute)

	srv.AddAppRouter(appRouter)
	log.Info("runninng server at localhost:", cfg.HTTPPort)
	err = srv.Run()

	if err != nil {
		log.Error("could not start the server.")
		log.Error("error running service.")
		time.Sleep(time.Millisecond * 50) // needed for printing all messages before exiting
		os.Exit(1)
	}
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
