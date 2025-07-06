package crypto

import (
	"encoding/asn1"
)

func BerDecode(data []byte) ([]asn1.RawValue, error) {
	var rawValues []asn1.RawValue
	_, err := asn1.Unmarshal(data, &rawValues)
	if err != nil {
		return nil, err
	}
	return rawValues, nil
}
