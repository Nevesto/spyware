package firefox

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func FindHistoryFile() (string, error) {
	profilePaths, err := GetFirefoxProfilePaths()
	if err != nil {
		return "", fmt.Errorf("failed to get Firefox profile paths: %w", err)
	}

	for _, profile := range profilePaths {
		historyFile := filepath.Join(profile, "places.sqlite")
		if _, err := os.Stat(historyFile); err == nil {
			return historyFile, nil
		}
	}

	return "", fmt.Errorf("Firefox History file not found")
}

func ReadHistory(historyFile string) ([]string, error) {
	db, err := sql.Open("sqlite", historyFile)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT url FROM moz_places")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urls []string
	for rows.Next() {
		var url string
		if err := rows.Scan(&url); err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	return urls, nil
}
