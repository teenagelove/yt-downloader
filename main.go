package main

import (
	"log"
	tgClient "yt-donwloader/clients/telegram"
	"yt-donwloader/consumer/event-consumer"
	"yt-donwloader/events/telegram"
)

const (
	tgBotHost = "api.telegram.org"
	batchSize = 100
)

func main() {
	//cfg := config.MustLoad()
	//test12.TestAudio()
	eventsProcessor := telegram.New(
		//tgClient.New(tgBotHost, cfg.TgBotToken),
		tgClient.New(tgBotHost, "6522972257:AAHqiMteOFyXw7xWuXI9kNZW2FzBQStGV-g"),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}
