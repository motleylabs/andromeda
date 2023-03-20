package request

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func ProcessPost(apiURL string, payload []byte) ([]byte, error) {
	apiKey := os.Getenv("API_KEY")

	c := &http.Client{}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", apiKey)

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("request error")
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}
