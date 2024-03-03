package config

import (
	"flag"
	"log"
)

// token TELEGRAM_BOT_TOKEN
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
