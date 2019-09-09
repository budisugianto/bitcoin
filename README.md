# Bitcoin

## Library for manipulating bitcoin addresses

### Usage

#### Generate a btc private/public key pair

```go
import btcpk "github.com/budisugianto/bitcoin/addr/private"

pk := btcpk.MustGenerate()

fmt.Printf("BTC Private key address (wallet format): %s", pk)
fmt.Printf("BTC                  Public key address: %s", pk.PublicKey())
```

#### Parse a btc public key address
```go
import btcpub "github.com/budisugianto/bitcoin/addr/public"

pubKey := btcpub.MustDecode("19Bq1gipWrLxFGqVH41Un2suWnGzWxNjbZ")

fmt.Printf("BTC public key: %s\n", pubKey)
fmt.Printf("  network: %s (%d, 0x%02X)\n", pubKey.Network, int(pubKey.Network), int(pubKey.Network))
fmt.Printf("  address: %v\n", pubKey.Address)
```

#### Parse a btc private key address in wallet format
```go
import btcpk "github.com/budisugianto/bitcoin/addr/private"

privKey := btcpk.MustDecode("5Ju6hf57BPdusMDUg4C6gPKiauXSahHVnTGmTNHoJeGUwJHeqSY")

fmt.Printf("BTC private key: %s\n", privKey)
fmt.Printf("    network: %s (%d, 0x%02X)\n", privKey.Network, int(privKey.Network), int(privKey.Network))
fmt.Printf("   exponent: %v\n", privKey.Exponent())
fmt.Printf(" public key: %s\n", privKey.PublicKey())
fmt.Printf("              X: %v\n", privKey.XBytes())
fmt.Printf("              Y: %v\n", privKey.YBytes())
```


## Testing

### Run test cases

`make`

### Run test cases and benchmarks

`make GOTESTOPT='-bench=.'`


## License

MIT

## Copyright

Copyright (C) 2013-2015 Barry Allard
