package main

import (
	"log"
	tgClient "yt-donwloader/clients/telegram"
	"yt-donwloader/config"
	"yt-donwloader/consumer/event-consumer"
	"yt-donwloader/events/telegram"
)

const (
	tgBotHost = "api.telegram.org"
	batchSize = 100
)

func main() {
	cfg := config.MustLoad()
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, cfg.TgBotToken),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}
