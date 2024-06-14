package telegram

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"yt-donwloader/lib/cleaner"
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

	newPath := "/bin/ffmpeg"

	// Добавляем новый путь в PATH.
	if err := converter.AddPath(newPath); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Путь успешно добавлен.")
	}

	// Проверяем значение PATH после добавления нового пути.
	fmt.Println("Текущее значение PATH:", os.Getenv("PATH"))

	output, err := converter.ExecuteCommand("ls /bin")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", output)
	}

	if isYoutube(text) {
		err := p.sendWait(ctx, chatID)
		if err != nil {
			return err
		}

		fileName, err := p.saveVideo(text)
		if err != nil {
			return p.tg.SendMessage(ctx, chatID, msgOops)
		}

		audioName, err := p.convertVideo(fileName)
		if err != nil {
			return p.tg.SendMessage(ctx, chatID, msgOops)
		}

		return p.sendAudio(ctx, chatID, audioName)
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
	fileName, err = downloader.Download(link)

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

func (p *Processor) sendWait(ctx context.Context, chatID int) error {
	return p.tg.SendMessage(ctx, chatID, msgWait)
}

func (p *Processor) sendAudio(ctx context.Context, chatID int, audioPath string) error {
	err := p.tg.SendAudio(ctx, chatID, audioPath)
	cleaner.ClearDirectory()
	return err
}

func isYoutube(text string) bool {
	return isURL(text)
}

func isURL(text string) bool {
	u, err := url.Parse(text)
	return err == nil && u.Host != ""
}
