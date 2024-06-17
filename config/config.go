package config

import (
	"flag"
	"log"
)

type Config struct {
	TgBotToken string
}

func MustLoad() Config {
	tgBotTokenToken := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *tgBotTokenToken == "" {
		log.Fatal("token is not specified")
	}

	return Config{
		TgBotToken: *tgBotTokenToken,
	}
}
