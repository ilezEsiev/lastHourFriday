package DB

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetAllChatIDsFromFile() ([]string, error) {
	filePath := "internal/DB/chatIDs.txt"

	// Открываем файл в режиме чтения
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer file.Close()

	var keys []string

	// Считываем значения из файла
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line != "" {
			keys = append(keys, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при чтении файла: %v", err)
	}

	return keys, nil
}
