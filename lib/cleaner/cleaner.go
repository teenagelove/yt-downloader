package cleaner

import (
	"log"
	"os"
	"path/filepath"
)

func clearDirectory(dirPath string) error {
	// Чтение всех файлов и подкаталогов в директории
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	// Удаление каждого файла и подкаталога
	for _, entry := range entries {
		filePath := filepath.Join(dirPath, entry.Name())

		err := os.RemoveAll(filePath)
		if err != nil {
			return err
		}
	}
	return nil
}

func ClearDirectory() {
	dirPath := "./storage/" // Укажите путь к директории

	err := clearDirectory(dirPath)
	if err != nil {
		log.Println("Error clearing directory:", err)
		return
	}

	log.Println("Directory cleared successfully")
}
