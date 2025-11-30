package input

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/kaugesaar/aoc/internal/auth"
)

// GetInput fetches input for a given year/day.
// It checks the cache first (inputs/{year}/{day:02d}.txt), and if missing,
// downloads from adventofcode.com using the session cookie.
func GetInput(year, day int) (io.Reader, error) {
	cachedPath := getCachePath(year, day)
	if data, err := os.ReadFile(cachedPath); err == nil {
		return bytes.NewReader(data), nil
	}

	if err := Download(year, day); err != nil {
		return nil, err
	}

	data, err := os.ReadFile(cachedPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read cached input: %w", err)
	}

	return bytes.NewReader(data), nil
}

// Download explicitly downloads input from adventofcode.com.
func Download(year, day int) error {
	session, err := auth.GetSession()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Cookie", fmt.Sprintf("session=%s", session))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to download input: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download input: HTTP %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	cachedPath := getCachePath(year, day)
	if err := os.MkdirAll(filepath.Dir(cachedPath), 0o755); err != nil {
		return fmt.Errorf("failed to create inputs directory: %w", err)
	}

	if err := os.WriteFile(cachedPath, data, 0o644); err != nil {
		return fmt.Errorf("failed to write cached input: %w", err)
	}

	return nil
}

func getCachePath(year, day int) string {
	return filepath.Join("inputs", fmt.Sprintf("%d", year), fmt.Sprintf("%02d.txt", day))
}
