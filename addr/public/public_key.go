package public

import (
	"fmt"
	"github.com/budisugianto/bitcoin/network"
	"github.com/budisugianto/bitcoin/util/key"
)

const AddressSize = 20 // bytes

type PublicKey struct {
	network.Network
	Address []byte
}

func NewFromNetworkAndAddress(net network.Network, address []byte) (pubKey *PublicKey, err error) {
	if len(address) != AddressSize {
		return nil, fmt.Errorf("Bitcoin address has incorrect number of byte(s), actual: %d, expected: %d", len(address), AddressSize)
	}
	return &PublicKey{
		Network: net,
		Address: address,
	}, nil
}

func NewFromAddress(address []byte) (pubKey *PublicKey, err error) {
	return NewFromNetworkAndAddress(network.Main, address)
}

func (pubKey PublicKey) PublicAddressPrefix() (addrPrefix byte, err error) {
	return pubKey.Network.PublicAddressPrefix()
}

func (pubKey PublicKey) Encode() (encoded string, err error) {
	addrPrefix, err := pubKey.PublicAddressPrefix()
	if err != nil {
		return
	}
	data := append([]byte{addrPrefix}, pubKey.Address...)
	return key.Encode(data), nil
}

func (pubKey PublicKey) MustEncode() (encoded string) {
	encoded, err := pubKey.Encode()
	if err != nil {
		panic(err)
	}
	return
}

func (pubKey PublicKey) String() string {
	return pubKey.MustEncode()
}

func Decode(encoded string) (pubKey *PublicKey, err error) {
	decoded, err := key.Decode(encoded)
	if err != nil {
		return
	}
	network, err := network.DecodePublicAddressPrefix(decoded[0])
	if err != nil {
		return
	}
	return NewFromNetworkAndAddress(network, decoded[1:])
}

func MustDecode(encoded string) (pubKey *PublicKey) {
	pubKey, err := Decode(encoded)
	if err != nil {
		panic(err)
	}
	return
}
