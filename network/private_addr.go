package network

import (
	"fmt"

	"github.com/budisugianto/bitcoin/util/inversemap"
)

var (
	privateNetworksToPrefixes = map[Network]byte{Main: 0x5A, Test: 0xEF}
	privatePrefixesToNetworks = inversemap.InverseMap(privateNetworksToPrefixes).(map[byte]Network)
)

func (net Network) PrivateAddressPrefix() (addrPrefix byte, err error) {
	a, ok := privateNetworksToPrefixes[net]
	if !ok {
		return 0, fmt.Errorf("%s", net)
	}
	return a, nil
}

func DecodePrivateAddressPrefix(addrPrefix byte) (net Network, err error) {
	n, ok := privatePrefixesToNetworks[addrPrefix]
	if !ok {
		return 0, fmt.Errorf("Unknown bitcoin private network address prefix %d (0x%02x)", addrPrefix, addrPrefix)
	}
	return n, nil
}
