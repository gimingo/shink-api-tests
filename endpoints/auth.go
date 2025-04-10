package endpoints

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	firebaseKey     = "fake-key"
	authEmulatorURL = "http://localhost:9099"
)

type signUpRequest struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type signUpResponse struct {
	IDToken      string     `json:"idToken"`
	Email        string     `json:"email"`
	RefreshToken string     `json:"refreshToken"`
	ExpiresIn    string     `json:"expiresIn"`
	LocalID      string     `json:"localId"`
	Error        *authError `json:"error,omitempty"`
}

type authError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Errors  []struct {
		Message string `json:"message"`
		Domain  string `json:"domain"`
		Reason  string `json:"reason"`
	} `json:"errors"`
}

func signUp(email, password string) (string, error) {
	reqBody := signUpRequest{
		Email:             email,
		Password:          password,
		ReturnSecureToken: true,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("%s/identitytoolkit.googleapis.com/v1/accounts:signUp?key=%s", authEmulatorURL, firebaseKey)
	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var res signUpResponse
	if err := json.Unmarshal(respBody, &res); err != nil {
		return "", err
	}
	if res.IDToken != "" {
		return res.IDToken, nil
	}
	if res.Error != nil {
		return "", fmt.Errorf("signUp error: %s", res.Error.Message)
	}
	return "", fmt.Errorf("unexpected signUp response: %s", string(respBody))
}

func signIn(email, password string) (string, error) {
	reqBody := signUpRequest{
		Email:             email,
		Password:          password,
		ReturnSecureToken: true,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("%s/identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s", authEmulatorURL, firebaseKey)
	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var res signUpResponse
	if err := json.Unmarshal(respBody, &res); err != nil {
		return "", err
	}
	if res.IDToken != "" {
		return res.IDToken, nil
	}
	if res.Error != nil {
		return "", fmt.Errorf("signIn error: %s", res.Error.Message)
	}
	return "", fmt.Errorf("unexpected signIn response: %s", string(respBody))
}

func generateRandomEmail() string {
	b := make([]byte, 5)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x@example.com", b)
}

func GetAuthToken() (string, error) {
	email := generateRandomEmail()
	password := "123qwe123"
	token, err := signUp(email, password)
	if err != nil && strings.Contains(err.Error(), "EMAIL_EXISTS") {
		return signIn(email, password)
	}
	return token, err
}
