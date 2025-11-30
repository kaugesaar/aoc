package auth

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const sessionFile = ".aoc/session"

// GetSession reads the session cookie from .aoc/session.
func GetSession() (string, error) {
	data, err := os.ReadFile(sessionFile)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("session not configured: run 'aoc login' first")
		}
		return "", fmt.Errorf("failed to read session file: %w", err)
	}
	return strings.TrimSpace(string(data)), nil
}

// SaveSession stores the session cookie to .aoc/session.
func SaveSession(cookie string) error {
	dir := filepath.Dir(sessionFile)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("failed to create .aoc directory: %w", err)
	}

	cookie = strings.TrimSpace(cookie)
	if err := os.WriteFile(sessionFile, []byte(cookie+"\n"), 0o600); err != nil {
		return fmt.Errorf("failed to write session file: %w", err)
	}

	return nil
}

func IsConfigured() bool {
	_, err := os.Stat(sessionFile)
	return err == nil
}
