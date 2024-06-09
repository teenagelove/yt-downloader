package cleaner

import (
	"io/ioutil"
	"log"
	"os"
)

func clearDirectory(dirPath string) error {
	// Чтение всех файлов и подкаталогов в директории
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	// Удаление каждого файла и подкаталога
	for _, file := range files {
		filePath := dirPath + "/" + file.Name()
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
