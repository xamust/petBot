package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/xamust/petbot/service_bot/internal/app/botapp"
	"log"
)

var configsPath string

func init() {
	flag.StringVar(&configsPath, "configs-path", "configs/config.toml", "Path to configs...")
}

func main() {
	flag.Parse()
	configs := botapp.NewConfig()
	meta, err := toml.DecodeFile(configsPath, configs)
	if err != nil {
		log.Fatalln(err)
	}
	if len(meta.Undecoded()) != 0 {
		log.Fatal("Undecoded configs param: ", meta.Undecoded())
	}
	//start bot service...
	botService := botapp.NewBot(configs)
	if err = botService.Start(); err != nil {
		log.Fatalln("Error on start:", err)
	}
}
