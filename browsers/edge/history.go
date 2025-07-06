package edge

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func FindHistoryFile() (string, error) {
	edgeUserDataPath, err := GetEdgeUserDataPath()
	if err != nil {
		return "", fmt.Errorf("failed to get Edge user data path: %w", err)
	}

	historyPath := filepath.Join(edgeUserDataPath, "History")
	if _, err := os.Stat(historyPath); err == nil {
		return historyPath, nil
	}

	profiles, err := GetEdgeProfilePaths()
	if err != nil {
		return "", fmt.Errorf("failed to get Edge profile paths: %w", err)
	}

	for _, profile := range profiles {
		historyPath = filepath.Join(profile, "History")
		if _, err := os.Stat(historyPath); err == nil {
			return historyPath, nil
		}
	}

	return "", fmt.Errorf("Edge History file not found")
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
