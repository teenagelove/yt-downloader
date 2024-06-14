package converter

import (
	"bytes"
	"os/exec"
)

func ExecuteCommand(command string) (string, error) {
	// Разделяем команду и её аргументы для выполнения.
	cmd := exec.Command(command)

	// Создаем буфер для хранения стандартного вывода.
	var out bytes.Buffer
	// Перенаправляем стандартный вывод команды в буфер.
	cmd.Stdout = &out

	// Выполняем команду и ждем её завершения.
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	// Возвращаем результат выполнения команды.
	return out.String(), nil
}
