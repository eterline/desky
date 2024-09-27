package notification

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Initialize API
func InitGotify(url, token string) (gotifyReq, error) {
	err := checkURL(url)
	if err != nil {
		return gotifyReq{}, err
	}
	return gotifyReq{
		uri:   url,
		token: token,
	}, nil
}

func (g *gotifyReq) Send(title, msg string, priority int) error {
	body := map[string]interface{}{
		"message":  msg,
		"priority": priority,
		"title":    title,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", g.uri, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+g.token)
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}
