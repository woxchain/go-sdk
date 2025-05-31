# Woxchain Go SDK

Official Go SDK for interacting with Woxchain Testnet.

## Usage

```go
import "github.com/woxchain/go-sdk/wox"

func main() {
    wox.Connect("https://rpc.woxscan.com")

    balance, _ := wox.GetBalance("0x1234...")
    fmt.Println("Balance:", balance)
}
```
