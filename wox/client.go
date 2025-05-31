package wox

import (
    "fmt"
    "net/http"
    "bytes"
    "io/ioutil"
)

var rpcURL string

func Connect(url string) {
    rpcURL = url
    fmt.Println("Connected to Woxchain RPC:", rpcURL)
}

func makeRPCRequest(method string, params []interface{}) ([]byte, error) {
    payload := fmt.Sprintf(`{"jsonrpc":"2.0","method":"%s","params":%v,"id":1}`, method, params)
    resp, err := http.Post(rpcURL, "application/json", bytes.NewBuffer([]byte(payload)))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    return ioutil.ReadAll(resp.Body)
}
