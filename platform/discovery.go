package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// ConfigDiscovery represents the JSON input configuration
type ConfigDiscovery struct {
	Token string `json:"token"`
}

type App struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Discover retrieves heroku user info
func Discover(token string) ([]App, error) {
	var apps []App

	url := "https://api.heroku.com/apps"

	client := http.DefaultClient

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request execution failed: %w", err)
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&apps); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return apps, nil
}

func HerokuIntegrationDiscovery(cfg ConfigDiscovery) ([]App, error) {
	// Check for the token
	if cfg.Token == "" {
		return nil, errors.New("token must be configured")
	}

	return Discover(cfg.Token)
}
