package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"errors"

	"golang.org/x/crypto/pbkdf2"
)

var (
	errBadCipher     = errors.New("bad cipher text")
	errCipherKey     = errors.New("invalid cipher key")
	errDecryptFailed = errors.New("decrypt failed")
)

func DPAPI(data []byte) ([]byte, error) {
	return dpapi(data)
}

func AesGCMDecrypt(data, key, nonce []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return gcm.Open(nil, nonce, data, nil)
}

func AesCBCDecrypt(data, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(data)%aes.BlockSize != 0 {
		return nil, errBadCipher
	}
	bm := cipher.NewCBCDecrypter(block, iv)
	bm.CryptBlocks(data, data)
	return data, nil
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

func DeriveKey(password, salt []byte) ([]byte, error) {
	if len(password) == 0 {
		return nil, errCipherKey
	}
	return pbkdf2.Key(password, salt, 1, 16, sha1.New), nil
}
