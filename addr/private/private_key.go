package private

import (
	"crypto/sha256"
	"math/big"

	"github.com/budisugianto/bitcoin/addr/public"
	"github.com/budisugianto/bitcoin/network"
	"github.com/budisugianto/bitcoin/util/key"
	"golang.org/x/crypto/ripemd160"
)

type PrivateKey struct {
	network.Network
	privateKey ECDSAPrivateKey
}

var publicAddrFirstRoundFirstByte = []byte{0x04}

// make sure pk can be encoded and decoded
func (pk PrivateKey) Encodable() (err error) {
	pkEncoded, err := pk.Encode()
	if err != nil {
		return
	}
	_, err = Decode(pkEncoded)
	return
}

func NewPrivateKey(net network.Network, ecdsaPk ECDSAPrivateKey) (pk *PrivateKey, err error) {
	_pk := &PrivateKey{
		Network:    net,
		privateKey: ecdsaPk,
	}
	err = _pk.Encodable()
	if err == nil {
		pk = _pk
	}
	return
}

func Generate() (pk *PrivateKey, err error) {
	ecdsaPk, err := NewECDSAPrivateKey()
	if err != nil {
		return
	}
	pk, err = NewPrivateKey(network.Main, ecdsaPk)
	return
}

func MustGenerate() (pk *PrivateKey) {
	pk, err := Generate()
	if err != nil {
		panic(err)
	}
	return
}

func NewFromNetworkAndExponent(network network.Network, exponent []byte) (pk *PrivateKey, err error) {
	ecdsa, err := NewECDSAPrivateKeyFromExponent(exponent)
	if err != nil {
		return
	}
	return &PrivateKey{
		Network:    network,
		privateKey: *ecdsa,
	}, nil
}

func NewFromExponent(exponent []byte) (pk *PrivateKey, err error) {
	return NewFromNetworkAndExponent(network.Main, exponent)
}

func (pk PrivateKey) PublicKey() *public.PublicKey {
	return &public.PublicKey{
		Network: pk.Network,
		Address: pk.PublicAddress(),
	}
}

func (pk PrivateKey) Exponent() []byte {
	return padToSize(pk.privateKey.priv, ExponentSize)
}

func padToSize(buf []byte, sz int) []byte {
	return append(make([]byte, sz-len(buf), sz), buf...)
}

func (pk PrivateKey) PublicAddressString() (pubAddr string, err error) {
	return pk.PublicKey().Encode()
}

func (pk PrivateKey) PublicAddressPrefix() (privAddrPrefix byte, err error) {
	return pk.Network.PublicAddressPrefix()
}

func (pk PrivateKey) PrivateAddressPrefix() (pubAddrPrefix byte, err error) {
	return pk.Network.PrivateAddressPrefix()
}

func (pk PrivateKey) X() *big.Int {
	return pk.privateKey.PublicKey.X
}

func (pk PrivateKey) Y() *big.Int {
	return pk.privateKey.PublicKey.Y
}

func (pk PrivateKey) XBytes() []byte {
	return pk.privateKey.PublicKey.XBytes()
}

func (pk PrivateKey) YBytes() []byte {
	return pk.privateKey.PublicKey.YBytes()
}

func (pk PrivateKey) XBytesPadded() []byte {
	return padToSize(pk.XBytes(), ExponentSize)
}

func (pk PrivateKey) YBytesPadded() []byte {
	return padToSize(pk.YBytes(), ExponentSize)
}

func (pk PrivateKey) PublicAddress() (publicAddr []byte) {
	firstRound := sha256.New()
	firstRound.Write(publicAddrFirstRoundFirstByte)
	firstRound.Write(pk.XBytesPadded())
	firstRound.Write(pk.YBytesPadded())

	secondRound := ripemd160.New()
	secondRound.Write(firstRound.Sum(nil))

	publicAddr = secondRound.Sum(nil)
	return
}

func (pk PrivateKey) Encode() (result string, err error) {
	addrPrefix, err := pk.PrivateAddressPrefix()
	if err != nil {
		return
	}
	data := append([]byte{addrPrefix}, pk.Exponent()...)

	return key.Encode(data), nil
}

func (pk PrivateKey) MustEncode() (result string) {
	result, err := pk.Encode()
	if err != nil {
		panic(err)
	}
	return
}

func (pk PrivateKey) String() string {
	return pk.MustEncode()
}

func Decode(encoded string) (pk *PrivateKey, err error) {
	decoded, err := key.Decode(encoded)
	if err != nil {
		return
	}

	network, err := network.DecodePrivateAddressPrefix(decoded[0]) // dipro
	if err != nil {
		return
	}

	exponent := decoded[1:]
	return NewFromNetworkAndExponent(network, exponent)
}

func MustDecode(encoded string) (pk *PrivateKey) {
	pk, err := Decode(encoded)
	if err != nil {
		panic(err)
	}
	return
}
