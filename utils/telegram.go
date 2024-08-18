package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type MessageRequest struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func SendTelegramMessage(message string) error {
	botToken := os.Getenv("BOT_TOKEN")
	chatID := os.Getenv("CHAT_ID")
	telegramApiUrl := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	messageRequest := MessageRequest{
		ChatID: chatID,
		Text:   message,
	}

	requestBody, err := json.Marshal(messageRequest)
	if err != nil {
		return fmt.Errorf("error marshalling request body: %w", err)
	}

	response, err := http.Post(telegramApiUrl, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("error sending message: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message, status code: %d", response.StatusCode)
	}

	fmt.Println("Message sent, status code:", response.StatusCode)
	return nil
}
