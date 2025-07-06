package crypto

import (
	"os"
	"path/filepath"
)

func InitNSS(profilePath string) error {
	nssPath := filepath.Join(profilePath, "nss")
	if _, err := os.Stat(nssPath); os.IsNotExist(err) {
		return err
	}
	return nil
}
