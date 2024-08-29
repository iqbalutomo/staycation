package helper

import (
	"bytes"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func FetchAPI(apiURL, method string, headers map[string]string, body []byte) (string, error) {
	req, err := http.NewRequest(method, apiURL, bytes.NewReader(body))
	if err != nil {
		return "", nil
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"message": "failed to fetch data",
		})
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}
