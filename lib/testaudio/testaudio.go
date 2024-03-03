package testAudio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

const (
	telegramAPIEndpoint = "https://api.telegram.org/bot%s/%s"
	botToken            = "6522972257:AAHqiMteOFyXw7xWuXI9kNZW2FzBQStGV-g"
	chatID              = "461945654"
)

func Test() {
	audioFilePath := "LSD.mp3"
	sendAudio(botToken, chatID, audioFilePath)
}

func sendAudio(token, chatID, audioFilePath string) {
	url := fmt.Sprintf(telegramAPIEndpoint, token, "sendAudio")

	// Open the audio file
	audioFile, err := os.Open(audioFilePath)
	if err != nil {
		fmt.Println("Error opening audio file:", err)
		return
	}
	defer audioFile.Close()

	// Create a multipart form
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Create form file field for audio
	part, err := writer.CreateFormFile("audio", audioFilePath)
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return
	}

	// Copy audio file content to the form field
	_, err = io.Copy(part, audioFile)
	if err != nil {
		fmt.Println("Error copying audio content:", err)
		return
	}

	// Add chat ID to the form
	_ = writer.WriteField("chat_id", chatID)

	// Close the multipart writer
	err = writer.Close()
	if err != nil {
		fmt.Println("Error closing multipart writer:", err)
		return
	}

	// Create HTTP POST request
	req, err := http.NewRequest("GET", url, &requestBody)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected response status:", resp.Status)
		return
	}

	// Read and print the response body
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decoding response body:", err)
		return
	}
	fmt.Println("Response:", result)
}
