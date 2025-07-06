package brave

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func FindHistoryFile() (string, error) {
	braveUserDataPath, err := GetBraveUserDataPath()
	if err != nil {
		return "", fmt.Errorf("failed to get Brave user data path: %w", err)
	}

	historyPath := filepath.Join(braveUserDataPath, "History")
	if _, err := os.Stat(historyPath); err == nil {
		return historyPath, nil
	}

	profiles, err := GetBraveProfilePaths()
	if err != nil {
		return "", fmt.Errorf("failed to get Brave profile paths: %w", err)
	}

	for _, profile := range profiles {
		historyPath = filepath.Join(profile, "History")
		if _, err := os.Stat(historyPath); err == nil {
			return historyPath, nil
		}
	}

	return "", fmt.Errorf("Brave History file not found")
}

func ReadHistory(historyFile string) ([]string, error) {
	db, err := sql.Open("sqlite", historyFile)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT url FROM urls")
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