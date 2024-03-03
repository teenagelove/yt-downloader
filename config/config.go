package config

import (
	"flag"
	"log"
)

// token 6522972257:AAHqiMteOFyXw7xWuXI9kNZW2FzBQStGV-g
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
