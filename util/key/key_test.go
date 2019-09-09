package key

import (
	"bytes"
	"github.com/budisugianto/bitcoin/test/hex"
	"testing"
)

var decodeTests = [][]byte{
	{},
	{0x00},
	{0x01},
	{0x02},
	{0xA5},
	{0x5A},
	{0x7F},
	{0x80},
	{0xFF},
	{0xFE},
	{0x00, 0x00},
	{0x01, 0x00},
	{0xFF, 0x00},
	{0xFF, 0xFF},
	{0xFF, 0xFF, 0xFF},
}

var encodeTests = []keyTestCase{
	{"1Wh4bh", "00"},
	{"112edB6q", "0000"},
	{"1J69wFcKYkABMsTyLyAxJmUnpRe5kiGWHg", "00bb73ad508e67b211ecaba7d3d592ab46983d50dc"},
	{"19Bq1gipWrLxFGqVH41Un2suWnGzWxNjbZ", "0059cd3e7f18a5f732e38b1202f191234721a921c1"},
	{"5Ju6hf57BPdusMDUg4C6gPKiauXSahHVnTGmTNHoJeGUwJHeqSY", "808e9f8f4f9ba55e4b8372b3e468e16fd7782f7686e473ed486269d4bc9195416d"},
}

type keyTestCase struct {
	encoded    string
	decodedHex string // apply MustDecode -> []byte
}

func (ktc keyTestCase) Decoded() []byte {
	return hex.MustDecode(ktc.decodedHex)
}

func TestDecode(t *testing.T) {
	for _, expected := range decodeTests {
		actual, err := Decode(Encode(expected))
		if err != nil {
			t.Errorf("Decode error! err: %v", err)
		} else if bytes.Compare(expected, actual) != 0 {
			t.Errorf("Decode incorrect\nexpected: %s\n  actual: %s", hex.String(expected), hex.String(actual))
		}
	}
	for _, tc := range encodeTests {
		expected := tc.Decoded()
		actual, err := Decode(tc.encoded)
		if err != nil {
			t.Errorf("Decode error! err: %v", err)
		} else if bytes.Compare(expected, actual) != 0 {
			t.Errorf("Decode incorrect\nexpected: %s\n  actual: %s", hex.String(expected), hex.String(actual))
		}
	}
}

func TestEncode(t *testing.T) {
	for _, tc := range encodeTests {
		expected := tc.encoded
		actual := Encode(tc.Decoded())
		if actual != expected {
			expectedHx := hex.String(tc.Decoded())
			actualBin, _ := Decode(tc.encoded)
			actualHx := hex.String(actualBin)
			t.Errorf("Encode failed\nexpected: %s (encoded)\nexpected: %s (decoded)\n  actual: %s (encoded)\n  actual: %s (decoded)", expected, expectedHx, actual, actualHx)
		}
	}
}
