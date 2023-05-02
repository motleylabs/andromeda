package request

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func ProcessGet(url string) ([]byte, error) {
	c := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", "data-provider")
	req.Header.Set("X-API-KEY", os.Getenv("SOLSNIPER_KEY"))

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("%d error: %s", res.StatusCode, string(body))
	}

	return body, nil
}

const USER_AGENT = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36"

func ProcessBrowserGet(url string) ([]byte, error) {
	c := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", USER_AGENT)

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("%d error: %s", res.StatusCode, string(body))
	}

	return body, nil
}
