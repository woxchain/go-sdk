package wox

import "math/big"

func ToWei(eth float64) *big.Int {
    wei := new(big.Float).Mul(big.NewFloat(eth), big.NewFloat(1e18))
    result := new(big.Int)
    wei.Int(result)
    return result
}
