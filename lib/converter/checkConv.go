package converter

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ExecuteCommand(command string) (string, error) {
	// Разделяем команду и её аргументы для выполнения.
	cmd := exec.Command("bash", "-c", command)

	// Создаем буфер для хранения стандартного вывода.
	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	// Выполняем команду и ждем её завершения.
	// Выполняем команду.
	err := cmd.Run()
	if err != nil {
		// Выводим расширенную информацию об ошибке.
		return "", fmt.Errorf("ошибка выполнения команды: %s, stderr: %s, err: %v", command, stderr.String(), err)
	}

	return out.String(), nil

	// Возвращаем результат выполнения команды.
	return out.String(), nil
}

func AddPath(newPath string) error {
	// Получаем текущее значение переменной PATH.
	currentPath := os.Getenv("PATH")

	// Проверяем, если новый путь уже присутствует в PATH.
	if strings.Contains(currentPath, newPath) {
		fmt.Println("Путь уже присутствует в PATH.")
		return nil
	}

	// Добавляем новый путь к текущему значению PATH.
	newPathValue := fmt.Sprintf("%s:%s", currentPath, newPath)

	// Устанавливаем новое значение переменной PATH.
	if err := os.Setenv("PATH", newPathValue); err != nil {
		return fmt.Errorf("ошибка при установке нового значения PATH: %v", err)
	}

	fmt.Println("Новый путь добавлен в PATH:", newPathValue)
	return nil
}
