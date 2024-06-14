package converter

import (
	"bytes"
	"fmt"
	"os/exec"
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
