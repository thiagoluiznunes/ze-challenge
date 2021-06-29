package main

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo"
	"github.com/thiagoluiznunes/ze-challenge/infra/config"
	"github.com/thiagoluiznunes/ze-challenge/server"
	"github.com/thiagoluiznunes/ze-challenge/server/router"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Couldn't read config file")
		return
	}
	_, err = json.Marshal(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	e := echo.New()
	srv := server.Instance(e, cfg)

	appRouter := router.New(srv.Echo, cfg, "ze-delivery")
	srv.AddAppRouter(appRouter)
}
