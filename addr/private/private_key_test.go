package private

import (
	"bytes"
	"fmt"
	"github.com/budisugianto/bitcoin/test/hex"
	"github.com/budisugianto/bitcoin/util/base58"
	"testing"
)

func testCase(t *testing.T, expected_encoded_pk string, expected_decoded_exponent []byte) {
	pk, err := NewFromExponent(expected_decoded_exponent)
	if err != nil {
		t.Errorf("NewFromAddress failed, err: %s", err)
		return
	}
	encoded_pk, err := pk.Encode()
	if err != nil {
		t.Errorf("PrivateKey Encode() failed, err: %s", err)
		return
	}

	if expected_encoded_pk != encoded_pk {
		t.Errorf("Encode failed:\nexpected: %s\nexpected: %v\n  actual: %s\n  actual: %v", expected_encoded_pk, base58.MustDecode(expected_encoded_pk), encoded_pk, base58.MustDecode(encoded_pk))
		return
	}

	pk, err = Decode(expected_encoded_pk)
	if err != nil {
		t.Errorf("Decode err: %s\nexpected_encoded_pk: %v", err, expected_encoded_pk)
		return
	}

	if bytes.Compare(pk.privateKey.priv, expected_decoded_exponent) != 0 {
		t.Errorf("Decode failed.\nexpected: %s\n  actual: %s", expected_decoded_exponent, pk.privateKey.priv)
		return
	}
}

func Test0(t *testing.T) {
	exponent := hex.MustDecode("594cfd670ac0453236816deec9ee2a4924c58e95fba992472846c1000e0adf80")
	pk := "5JVcebpmNXH4eekGXpYLL526G8x7fdm9yt6TtxsCUkyXWpkM58j"
	testCase(t, pk, exponent)
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Generate()
		if err != nil {
			b.Errorf(fmt.Sprint("Could not generate bitcoin private key, err =", err))
		}
	}
}

func BenchmarkSetKey(b *testing.B) {
	exponent := hex.MustDecode("594cfd670ac0453236816deec9ee2a4924c58e95fba992472846c1000e0adf80")
	for i := 0; i < b.N; i++ {
		private_key, err := NewFromExponent(exponent)
		if err != nil || private_key == nil {
			b.Errorf(fmt.Sprint("Could not generate bitcoin private key, err =", err))
		}
	}
}
