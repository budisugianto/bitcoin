package key

import (
	"crypto/subtle"
	"errors"
	"github.com/budisugianto/bitcoin/util/base58"
	"github.com/budisugianto/bitcoin/util/checksum"
)

var checksumFailure = errors.New("Checksum failure")

func Decode(encoded string) (decoded []byte, err error) {
	decAndChksum, err := base58.Decode(encoded)
	if err != nil {
		return
	}

	// verify checksum
	dec := decAndChksum[:len(decAndChksum)-4]
	actualChecksum := checksum.Checksum(dec)
	expectedChecksum := decAndChksum[len(decAndChksum)-4:]
	if subtle.ConstantTimeCompare(actualChecksum, expectedChecksum) != 1 {
		return nil, checksumFailure
	}

	return dec, nil
}

func Encode(decoded []byte) string {
	return base58.Encode(append(decoded, checksum.Checksum(decoded)...))
}
