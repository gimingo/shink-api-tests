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

type CreateShinkResponse struct {
	ID               string   `json:"id"`
	UserID           string   `json:"user_id"`
	Name             string   `json:"name"`
	Hash             string   `json:"hash"`
	URL              string   `json:"url"`
	ConsumptionCount int      `json:"consumption_count"`
	Labels           []string `json:"labels"`
	ExpiresAt        string   `json:"expires_at"`
	CreatedAt        string   `json:"created_at"`
}

func CreateShink(token, name, target, hash string) (CreateShinkResponse, error) {
	urlStr := "http://localhost:8181/shink"
	expiresAt := time.Now().UTC().Add(72 * time.Hour).Format(time.RFC3339)
	data := url.Values{}
	data.Set("name", name)
	data.Set("target", target)
	data.Set("hash", hash)
	data.Set("expires_at", expiresAt)
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	if err != nil {
		return CreateShinkResponse{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Auth-Token", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return CreateShinkResponse{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return CreateShinkResponse{}, fmt.Errorf("non-OK HTTP status: %d. Response: %s", resp.StatusCode, string(body))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return CreateShinkResponse{}, err
	}
	var responsePayload CreateShinkResponse
	err = json.Unmarshal(body, &responsePayload)
	if err != nil {
		return CreateShinkResponse{}, err
	}
	return responsePayload, nil
}
