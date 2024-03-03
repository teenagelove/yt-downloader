package telegram

import (
	"context"
	"log"
	"net/url"
	"strings"
	"yt-donwloader/lib/converter"
	"yt-donwloader/lib/downloader"
	"yt-donwloader/lib/e"
	"yt-donwloader/lib/testaudio"
)

const (
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (p *Processor) doCmd(ctx context.Context, text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s", text, username)
	if isYoutube(text) {
		//fileName := p.saveVideo(ctx, chatID, text)
		//p.saveVideo(ctx, chatID, text)
		testAudio.Test()
		//return p.convertVideo(ctx, chatID, fileName)
		//return p.sendAudio(ctx, chatID)
	}

	switch text {
	case HelpCmd:
		return p.sendHelp(ctx, chatID)
	case StartCmd:
		return p.sendHello(ctx, chatID)
	default:
		return p.tg.SendMessage(ctx, chatID, msgUnknownCommand)
	}
}

func (p *Processor) saveVideo(ctx context.Context, chatID int, link string) (fileName string) {
	//defer func() { err = e.WrapIfErr("can't do command: save page", err) }()
	fileName = downloader.Donwload(link)

	return fileName
}

func (p *Processor) convertVideo(ctx context.Context, chatID int, fileName string) (err error) {
	defer func() { err = e.WrapIfErr("can't do command: save page", err) }()
	converter.Converter(fileName)

	return nil
}

func (p *Processor) sendHelp(ctx context.Context, chatID int) error {
	return p.tg.SendMessage(ctx, chatID, msgHelp)
}

func (p *Processor) sendHello(ctx context.Context, chatID int) error {
	return p.tg.SendMessage(ctx, chatID, msgHello)
}

func (p *Processor) sendAudio(ctx context.Context, chatID int) error {
	//audio, _ := os.Open("./storage/LSD - Audio (Official Video) ft. Sia, Diplo, Labrinth.mp4.mp3")
	return p.tg.SendAudio(ctx, chatID, "LSD.mp3")
	//return p.tg.SendMessage(ctx, chatID, msgVideo)
}

func isYoutube(text string) bool {
	return isURL(text)
}

func isURL(text string) bool {
	u, err := url.Parse(text)

	return err == nil && u.Host != ""
}
