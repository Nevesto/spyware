package chrome

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func GetChromeUserDataPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}
	return filepath.Join(homeDir, "AppData", "Local", "Google", "Chrome", "User Data"), nil
}

func GetChromeLoginDataPath() (string, error) {
	userDataPath, err := GetChromeUserDataPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(userDataPath, "Default", "Login Data"), nil
}

func GetChromeProfilePaths() ([]string, error) {
	userDataPath, err := GetChromeUserDataPath()
	if err != nil {
		return nil, err
	}

	var profiles []string

	profiles = append(profiles, filepath.Join(userDataPath, "Default"))

	profilePattern := filepath.Join(userDataPath, "Profile *")
	matches, err := filepath.Glob(profilePattern)
	if err != nil {
		return nil, fmt.Errorf("failed to glob profile directories: %w", err)
	}

	for _, match := range matches {
		info, err := os.Stat(match)
		if err != nil {
			continue
		}
		if info.IsDir() {
			profiles = append(profiles, match)
		}
	}

	return profiles, nil
}

func CopyFile(srcPath string) (string, error) {
	tmpfile, err := os.CreateTemp("", "temp_db")
	if err != nil {
		return "", err
	}
	defer tmpfile.Close()

	src, err := os.Open(srcPath)
	if err != nil {
		return "", err
	}
	defer src.Close()

	_, err = io.Copy(tmpfile, src)
	if err != nil {
		return "", err
	}

	return tmpfile.Name(), nil
}
