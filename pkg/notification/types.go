package notification

import (
	"errors"
	"net/url"
)

var errURI = errors.New("Uncorrect API URI")

type (
	telegramReq struct {
		chatID   int
		tokenBot string
	}
	gotifyReq struct {
		uri   string
		token string
	}
)

func checkURL(link string) error {
	parsedURL, err := url.Parse(link)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return errURI
	}
	return nil
}
