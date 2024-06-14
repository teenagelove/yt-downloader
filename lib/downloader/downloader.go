package downloader

import (
	"fmt"
	"github.com/kkdai/youtube/v2"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

// sanitizeFileName удаляет недопустимые символы из имени файла
func sanitizeFileName(fileName string) string {
	// Определите недопустимые символы
	_ = `\/:*?"<>|`
	replacer := strings.NewReplacer(
		"/", "_",
		"\\", "_",
		":", "_",
		"*", "_",
		"?", "_",
		"\"", "_",
		"<", "_",
		">", "_",
		"|", "_",
		//" ", "_", // Пробелы заменяем на подчеркивания
	)
	sanitizedFileName := replacer.Replace(fileName)

	// Ограничим длину имени файла до 200 символов
	maxFileNameLength := 200
	if utf8.RuneCountInString(sanitizedFileName) > maxFileNameLength {
		runes := []rune(sanitizedFileName)
		sanitizedFileName = string(runes[:maxFileNameLength])
	}

	return sanitizedFileName
}

func Download(link string) (outputFile string, err error) {
	// Проверяем наличие директории storage
	storageDir := "./storage"
	if _, err := os.Stat(storageDir); os.IsNotExist(err) {
		err := os.MkdirAll(storageDir, os.ModePerm)
		if err != nil {
			log.Printf("Error creating storage directory: %v\n", err)
			return "", err
		}
	}

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
	defer func(stream io.ReadCloser) {
		err := stream.Close()
		if err != nil {
			log.Printf("Error closing stream: %v\n", err)
		}
	}(stream)

	// Создаем файл для сохранения видео
	sanitizedTitle := sanitizeFileName(video.Title)
	fileName := fmt.Sprintf("%s/%s", storageDir, sanitizedTitle)
	file, err := os.Create(fileName + ".mp4")
	if err != nil {
		// Логируем ошибку и возвращаем её
		log.Printf("Error creating file: %v\n", err)
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("Error closing file: %v\n", err)
		}
	}(file)

	// Копируем поток в файл
	_, err = io.Copy(file, stream)
	if err != nil {
		// Логируем ошибку и возвращаем её
		log.Printf("Error copying stream to file: %v\n", err)
		return "", err
	}

	return fileName, nil
}
