package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Initialize API. "chat" <-- target user chat ID
func InitTg(chat int, token string) telegramReq {
	return telegramReq{
		chatID:   chat,
		tokenBot: token,
	}
}

func (t *telegramReq) Send(msg string) error {
	body := map[string]interface{}{
		"chat_id": t.chatID,
		"text":    msg,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}
	_, err = http.Post(
		fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.tokenBot),
		"application/json", bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return err
	}
	return nil
}
