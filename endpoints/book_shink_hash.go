package endpoints

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type BookShinkHashResponse struct {
	ShinkHash     string `json:"shink_hash"`
	BookingExpiry string `json:"booking_expiry"`
}

func BookShinkHash(token string) (string, error) {
	urlStr := "http://localhost:8181/shink/book-hash"
	expiry := time.Now().Add(72 * time.Hour).Format(time.RFC3339)

	data := url.Values{}
	data.Set("shink_expiry", expiry)

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Auth-Token", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("non-OK HTTP status: %d. Response: %s", resp.StatusCode, string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var responsePayload BookShinkHashResponse
	err = json.Unmarshal(body, &responsePayload)
	if err != nil {
		return "", err
	}

	return responsePayload.ShinkHash, nil
}
