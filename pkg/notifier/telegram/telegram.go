package telegram

import (
	"net/http"
	"os"
)

// Notify sends a notification by Telegram chat
func Notify(text string) {
	url := "https://api.telegram.org/bot" +
		os.Getenv("TELEGRAM_API_KEY") +
		"/sendMessage?chat_id=" +
		os.Getenv("TELEGRAM_CHAT_ID") +
		"&text=" +
		text

	_, _ = http.Get(url)
}
