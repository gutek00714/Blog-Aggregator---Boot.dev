package main

import (
	"fmt"

	"github.com/gutek00714/Blog-Aggregator---Boot.dev/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = cfg.SetUser("Gustaw")
	if err != nil {
		fmt.Println(err)
		return
	}
	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cfg)
}
