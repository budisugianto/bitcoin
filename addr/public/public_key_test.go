package public

import (
	"bytes"
	"github.com/budisugianto/bitcoin/test/hex"
	"testing"
)

// 20 + 4
// 0123456789012345678901234567890123456789
//           1         2         3
// "bb73ad508e67b211ecaba7d3d592ab46983d50dc 33164a23"
// "bb73ad508e67b211ecaba7d3d592ab46983d50dc"

func testCase(t *testing.T, expected_encoded_addr string, expected_decoded_addr []byte) {
	address, err := NewFromAddress(expected_decoded_addr)
	if err != nil {
		t.Errorf("NewFromAddress failed, err: %v", err)
		return
	}
	encoded_addr, err := address.Encode()
	if err != nil {
		t.Errorf("PublicKey Encode() failed, err: %v", err)
		return
	}

	if expected_encoded_addr != encoded_addr {
		t.Errorf("Encode failed.\nexpected: %s\n  actual: %s", expected_encoded_addr, encoded_addr)
		return
	}

	decoded_addr, err := Decode(expected_encoded_addr)
	if nil != err {
		t.Errorf("Decode error %s", err)
		return
	}

	if bytes.Compare(decoded_addr.Address, expected_decoded_addr) != 0 {
		t.Errorf("Decode failed.\n expected: %s\n  actual: %s", expected_decoded_addr, decoded_addr.Address)
		return
	}
}

func Test0(t *testing.T) {
	expected_decoded_addr := hex.MustDecode("bb73ad508e67b211ecaba7d3d592ab46983d50dc")
	expected_encoded_addr := "1J69wFcKYkABMsTyLyAxJmUnpRe5kiGWHg"
	testCase(t, expected_encoded_addr, expected_decoded_addr)
}
