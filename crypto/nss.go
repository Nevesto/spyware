package crypto

import (
	"bytes"
	"encoding/asn1"
	"errors"
)

var (
	ErrUnsupported = errors.New("unsupported")
)

type PBE struct {
	Data               []byte
	EntrySalt          []byte
	PasswordCheck      []byte
	Algorithm          asn1.ObjectIdentifier
	PasswordIterations int
}

type ASN1PBE struct {
	Algorithm  asn1.ObjectIdentifier
	Parameters struct {
		EntrySalt          []byte
		PasswordIterations int
	}
}

type ASN1PBEData struct {
	Data          []byte
	PasswordCheck []byte
}

func UnmarshalPBE(data []byte) (*PBE, error) {
	var pbeData ASN1PBEData
	_, err := asn1.Unmarshal(data, &pbeData)
	if err != nil {
		return nil, err
	}

	var pbe ASN1PBE
	_, err = asn1.Unmarshal(pbeData.Data, &pbe)
	if err != nil {
		return nil, err
	}

	return &PBE{
		Data:               pbeData.Data,
		EntrySalt:          pbe.Parameters.EntrySalt,
		PasswordCheck:      pbeData.PasswordCheck,
		Algorithm:          pbe.Algorithm,
		PasswordIterations: pbe.Parameters.PasswordIterations,
	}, nil
}

func (p *PBE) Decrypt(password []byte) ([]byte, error) {
	key, err := DeriveKey(password, p.EntrySalt)
	if err != nil {
		return nil, err
	}

	iv := bytes.Repeat([]byte{0x00}, 16)
	decrypted, err := AesCBCDecrypt(p.Data, key, iv)
	if err != nil {
		return nil, err
	}

	if !bytes.Equal(p.PasswordCheck, decrypted[len(decrypted)-len(p.PasswordCheck):]) {
		return nil, ErrUnsupported
	}

	return decrypted, nil
}

type Nss struct {
	Data []byte
}

func (n *Nss) Decrypt(password []byte) ([]byte, error) {
	return nil, ErrUnsupported
}

func NewNss(data []byte) (*Nss, error) {
	return &Nss{Data: data}, nil
}
