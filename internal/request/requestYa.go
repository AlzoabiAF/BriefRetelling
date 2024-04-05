package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const token = "y0_AgAAAABMYpAoAAoX4wAAAAD_h0nnAABXpW-lgLdA7b7igdbfAwIvkh9ykw"
const endpoint = "https://300.ya.ru/api/sharing-url"

type requestPost struct {
	ArticleUrl string `json:"article_url"`
}

type responsePost struct {
	Status     string `json:"status"`
	SharingUrl string `json:"sharing_url"`
}

func RequestMethod(url string) (string, error) {
	const op = "internal.request.RequestMethod"

	data := requestPost{ArticleUrl: url}

	body, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("%s: cannot to create json format: %w", op, err)
	}

	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(body))
	req.Header.Add("Authorization", fmt.Sprintf("OAuth %s", token))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("%s: failed to make a request: %w", op, err)
	}
	defer resp.Body.Close()

	var res responsePost
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return "", fmt.Errorf("%s: cannot to decode response body: %w", op, err)
	}

	if res.Status != "success" {
		return "", fmt.Errorf("%s: resp.Status is not successful: %s", op, res.Status)
	}

	return res.SharingUrl, nil
}
