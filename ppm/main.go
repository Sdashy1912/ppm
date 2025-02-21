package main

import (
	"log"

	"ppm/libnvth/api"
	"ppm/libnvth/configs"
)

func main() {
	err := configs.InitConfig()
	if err != nil {
		log.Fatalf("Could not read configurations: %s", err)
	}
	server, err := api.NewServer()
	if err != nil {
		log.Fatal("Could not initialize http server: ", err.Error())
	}
	server.Start()
	// logger := logging.NewLogger()
	// session, err := database.New()
	// if err != nil {
	// 	logger.WithField("module", "database").Error(err)
	// 	return
	// }
	// d := stats.NewStatsDAO(session)
	// d.Stats()
}
