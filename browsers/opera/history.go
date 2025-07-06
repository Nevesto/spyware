package opera

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func FindHistoryFile() (string, error) {
	operaUserDataPath, err := GetOperaUserDataPath()
	if err != nil {
		return "", fmt.Errorf("failed to get Opera user data path: %w", err)
	}

	historyFile := filepath.Join(operaUserDataPath, "History")
	if _, err := os.Stat(historyFile); err == nil {
		return historyFile, nil
	}

	return "", fmt.Errorf("Opera History file not found")
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
