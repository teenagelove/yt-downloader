package telegram

import (
	"context"
	"log"
	"net/url"
	"strings"
	"yt-donwloader/lib/converter"
	"yt-donwloader/lib/downloader"
	"yt-donwloader/lib/e"
)

const (
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (p *Processor) doCmd(ctx context.Context, text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s", text, username)
	if isYoutube(text) {
		fileName, _ := p.saveVideo(text)
		audioName, _ := p.convertVideo(fileName)
		return p.sendAudio(ctx, chatID, audioName)
		//return p.sendAudio(ctx, chatID, "./storage/ПОШЛАЯ МОЛЛИ, HOFMANNITA – #HABIBATI.mp4")
		//testAudio.Test()
		//return p.convertVideo(ctx, chatID, fileName)
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

func (p *Processor) saveVideo(link string) (fileName string, err error) {
	defer func() { err = e.WrapIfErr("can't do command: save video", err) }()
	fileName, err = downloader.Donwload(link)

	return fileName, err
}

func (p *Processor) convertVideo(fileName string) (outputFile string, err error) {
	defer func() { err = e.WrapIfErr("can't do command: convert video", err) }()
	outputFile, err = converter.Converter(fileName)

	return outputFile, err
}

func (p *Processor) sendHelp(ctx context.Context, chatID int) error {
	return p.tg.SendMessage(ctx, chatID, msgHelp)
}

func (p *Processor) sendHello(ctx context.Context, chatID int) error {
	return p.tg.SendMessage(ctx, chatID, msgHello)
}

func (p *Processor) sendAudio(ctx context.Context, chatID int, audioPath string) error {
	return p.tg.SendAudio(ctx, chatID, audioPath)
}

func isYoutube(text string) bool {
	return isURL(text)
}

func isURL(text string) bool {
	u, err := url.Parse(text)

	return err == nil && u.Host != ""
}
