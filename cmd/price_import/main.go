package main

import (
	"flag"
	"glancer/instance"
	"glancer/persistence"
	"glancer/usecase"
	"log"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "Config")
	flag.Parse()

	t := instance.NewInstance()
	err := t.Init(configPath, "")
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	pd := persistence.NewPriceDataPersistence(t)

	uc := usecase.NewPriceDataUseCase(
		pd,
	)

	log.Println(uc)

}
