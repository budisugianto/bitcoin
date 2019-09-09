package checksum

import (
	"bytes"
	"github.com/budisugianto/bitcoin/test/hex"
	"testing"
)

var (
	data1 = hex.MustDecode("00bb73ad508e67b211ecaba7d3d592ab46983d50dc")
	data2 = hex.MustDecode("01bb73ad508e67b211ecaba7d3d592ab46983d50dc")
	data3 = hex.MustDecode("00bb73ad508e67b211ecaba7d3d592ab46983d50de")
	data4 = hex.MustDecode("00bb73ad508e67b211ecaba7d3d592ab46983d50")
	data5 = hex.MustDecode("")
)

var (
	data0     = hex.MustDecode("00bb73ad508e67b211ecaba7d3d592ab46983d50dc")
	checksum0 = hex.MustDecode("33164a23")
)

func testCase(t *testing.T, data []byte, expected_checksum []byte) {
	checksum := Checksum(data)
	if 0 != bytes.Compare(checksum, expected_checksum) {
		t.Errorf("Checksum testcase failed:\nexpected: %s\n  actual: %s", expected_checksum, checksum)
		return
	}
}

func TestChecksum0(t *testing.T) {
	testCase(t, data0, checksum0)
}

func testCompare(t *testing.T, a, b []byte, expected bool) {
	if expected != Compare(a, b) {
		t.Errorf("Should have failed")
		return
	}
}

func testCompareSuccess(t *testing.T, a, b []byte) {
	testCompare(t, a, b, true)
}

func testCompareFail(t *testing.T, a, b []byte) {
	testCompare(t, a, b, false)
}

func TestCompare(t *testing.T) {

	testCompareFail(t, data1, data2)
	testCompareFail(t, data1, data3)
	testCompareFail(t, data1, data4)
	testCompareFail(t, data1, data5)
	testCompareSuccess(t, data1, data1)
	testCompareSuccess(t, data5, data5)
}
