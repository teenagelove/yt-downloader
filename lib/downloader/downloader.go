package downloader

import (
	"fmt"
	"github.com/kkdai/youtube/v2"
	"io"
	"log"
	"os"
)

func Download(link string) (outputFile string, err error) {
	videoID := link
	client := youtube.Client{}

	// Получаем видео по ID
	video, err := client.GetVideo(videoID)
	if err != nil {
		// Логируем ошибку и возвращаем её
		log.Printf("Error getting video: %v\n", err)
		return "", err
	}

	// Фильтруем форматы видео с аудиоканалами
	formats := video.Formats.WithAudioChannels()
	if len(formats) == 0 {
		err := fmt.Errorf("no formats with audio channels available")
		log.Printf("Error: %v\n", err)
		return "", err
	}

	// Получаем поток видео
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		// Логируем ошибку и возвращаем её
		log.Printf("Error getting stream: %v\n", err)
		return "", err
	}
	defer stream.Close()

	// Создаем файл для сохранения видео
	fileName := "./storage/" + video.Title
	file, err := os.Create(fileName + ".mp4")
	if err != nil {
		// Логируем ошибку и возвращаем её
		log.Printf("Error creating file: %v\n", err)
		return "", err
	}
	defer file.Close()

	// Копируем поток в файл
	_, err = io.Copy(file, stream)
	if err != nil {
		// Логируем ошибку и возвращаем её
		log.Printf("Error copying stream to file: %v\n", err)
		return "", err
	}

	return fileName, nil
}
