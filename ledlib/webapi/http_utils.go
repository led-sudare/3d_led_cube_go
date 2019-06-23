package webapi

import (
	"bytes"
	"net/http"
	"time"
)

func HttpJsonPost(url string, body []byte) error {

	req, err := http.NewRequest("POST",
		url,
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Duration(100 * time.Millisecond),
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}
