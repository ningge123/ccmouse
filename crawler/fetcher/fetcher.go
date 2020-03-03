package fetcher

import (
	"fmt"
	"io"
	"net/http"
)

func Fetch(url string) (io.ReadCloser, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")
	response, err := client.Do(request)

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("worng status code: %d", response.StatusCode)
	}

	return response.Body, nil
}
