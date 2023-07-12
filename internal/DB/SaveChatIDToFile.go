package DB

import (
	"bufio"
	"fmt"
	"os"
)

func SaveChatIDToFile(chatID int64) error {
	filePath := "internal/DB/chatIDs.txt"

	// Открываем файл в режиме записи (создаем, если не существует)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer file.Close()

	// Записываем chat ID в файл
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(fmt.Sprintf("%d\n", chatID))
	if err != nil {
		return fmt.Errorf("не удалось записать chat ID в файл: %v", err)
	}

	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("не удалось записать данные в файл: %v", err)
	}

	return nil
}
