package wox

import (
    "encoding/json"
    "math/big"
)

type BalanceResult struct {
    Result string `json:"result"`
}

func GetBalance(address string) (*big.Int, error) {
    data, err := makeRPCRequest("eth_getBalance", []interface{}{address, "latest"})
    if err != nil {
        return nil, err
    }

    var result BalanceResult
    err = json.Unmarshal(data, &result)
    if err != nil {
        return nil, err
    }

    balance := new(big.Int)
    balance.SetString(result.Result[2:], 16)
    return balance, nil
}
