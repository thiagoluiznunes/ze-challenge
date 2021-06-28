package main

import (
	"encoding/json"
	"fmt"

	"github.com/thiagoluiznunes/ze-challenge/infra/config"
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

}
