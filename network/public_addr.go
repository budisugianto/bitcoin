package network

import (
	"fmt"
	"github.com/budisugianto/bitcoin/util/inversemap"
)

var (
	publicNetworksToPrefixes = map[Network]byte{Main: 0x00, Test: 0x6F}
	publicPrefixesToNetworks = inversemap.InverseMap(publicNetworksToPrefixes).(map[byte]Network)
)

func (net Network) PublicAddressPrefix() (addrPrefix byte, err error) {
	a, ok := publicNetworksToPrefixes[net]
	if !ok {
		return 0, fmt.Errorf("%s", net)
	}
	return a, nil
}

func DecodePublicAddressPrefix(addrPrefix byte) (net Network, err error) {
	n, ok := publicPrefixesToNetworks[addrPrefix]
	if !ok {
		return 0, fmt.Errorf("Unknown bitcoin public address prefix %d (0x%02X)", addrPrefix, addrPrefix)
	}
	return n, nil
}
