package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// VideoMapping stores the alias to URL mappings
type VideoMapping struct {
	Aliases map[string]string `json:"aliases"`
}

// HistoryEntry represents a single playback history entry
type HistoryEntry struct {
	Alias     string    `json:"alias"`
	URL       string    `json:"url"`
	PlayedAt  time.Time `json:"played_at"`
	VideoMode bool      `json:"video_mode"`
}

// PlaybackHistory stores recent playback history
type PlaybackHistory struct {
	Recent []HistoryEntry `json:"recent"`
}

// getConfigPath returns the path to the config file
func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}

	configDir := filepath.Join(homeDir, ".config", "clitube")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create config directory: %w", err)
	}

	return filepath.Join(configDir, "videos.json"), nil
}

// LoadMappings loads video mappings from the config file
func LoadMappings() (*VideoMapping, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	// If file doesn't exist, return empty mappings
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return &VideoMapping{Aliases: make(map[string]string)}, nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var mapping VideoMapping
	if err := json.Unmarshal(data, &mapping); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	if mapping.Aliases == nil {
		mapping.Aliases = make(map[string]string)
	}

	return &mapping, nil
}

// SaveMappings saves video mappings to the config file
func SaveMappings(mapping *VideoMapping) error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(mapping, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal mappings: %w", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

// AddMapping adds a new alias to URL mapping
func AddMapping(alias, url string) error {
	mappings, err := LoadMappings()
	if err != nil {
		return err
	}

	mappings.Aliases[alias] = url
	return SaveMappings(mappings)
}

// GetURL retrieves the URL for a given alias
func GetURL(alias string) (string, error) {
	mappings, err := LoadMappings()
	if err != nil {
		return "", err
	}

	url, exists := mappings.Aliases[alias]
	if !exists {
		return "", fmt.Errorf("alias '%s' not found", alias)
	}

	return url, nil
}

// getHistoryPath returns the path to the history file
func getHistoryPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}

	configDir := filepath.Join(homeDir, ".config", "clitube")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create config directory: %w", err)
	}

	return filepath.Join(configDir, "history.json"), nil
}

// LoadHistory loads playback history from the config file
func LoadHistory() (*PlaybackHistory, error) {
	historyPath, err := getHistoryPath()
	if err != nil {
		return nil, err
	}

	// If file doesn't exist, return empty history
	if _, err := os.Stat(historyPath); os.IsNotExist(err) {
		return &PlaybackHistory{Recent: []HistoryEntry{}}, nil
	}

	data, err := os.ReadFile(historyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read history file: %w", err)
	}

	var history PlaybackHistory
	if err := json.Unmarshal(data, &history); err != nil {
		return nil, fmt.Errorf("failed to parse history file: %w", err)
	}

	if history.Recent == nil {
		history.Recent = []HistoryEntry{}
	}

	return &history, nil
}

// SaveHistory saves playback history to the config file
func SaveHistory(history *PlaybackHistory) error {
	historyPath, err := getHistoryPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal history: %w", err)
	}

	if err := os.WriteFile(historyPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write history file: %w", err)
	}

	return nil
}

// AddToHistory adds a playback entry to history, keeping only the last 3
func AddToHistory(alias, url string, videoMode bool) error {
	history, err := LoadHistory()
	if err != nil {
		return err
	}

	// Create new entry
	entry := HistoryEntry{
		Alias:     alias,
		URL:       url,
		PlayedAt:  time.Now(),
		VideoMode: videoMode,
	}

	// Add to beginning of list
	history.Recent = append([]HistoryEntry{entry}, history.Recent...)

	// Keep only last 3 entries
	if len(history.Recent) > 3 {
		history.Recent = history.Recent[:3]
	}

	return SaveHistory(history)
}

// GetRecentHistory returns the recent playback history
func GetRecentHistory() ([]HistoryEntry, error) {
	history, err := LoadHistory()
	if err != nil {
		return nil, err
	}

	return history.Recent, nil
}

// IsFirstRun checks if this is the first time running the app
func IsFirstRun() (bool, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return false, err
	}

	historyPath, err := getHistoryPath()
	if err != nil {
		return false, err
	}

	// Check if neither config file exists
	_, configErr := os.Stat(configPath)
	_, historyErr := os.Stat(historyPath)

	return os.IsNotExist(configErr) && os.IsNotExist(historyErr), nil
}
