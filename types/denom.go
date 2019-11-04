package types

import (
	"math/big"
	"strconv"
)

const (
	UndDenom  = "und"  // 1 (base unit)
	NundDenom = "nund" // 10^-12 (pico)

	UndPow  = 1e9  // multiplier for converting from und to (nano) nund
	NundPow = 1e-9 // multiplier for converting from (nano) nund to und
)

func ConvertUndDenomination(amount string, from string, to string) (string, error) {

	if from == to {
		return amount + from, nil
	}

	switch from {
	case UndDenom: // from und to pund
		fromAmt, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			return "", err
		}
		fromAmtBf := new(big.Float).SetFloat64(fromAmt)
		res := fromAmtBf.Mul(fromAmtBf, big.NewFloat(UndPow))
		result := new(big.Int)
		res.Int(result)
		return result.String() + to, nil
	case NundDenom: // from pund to und
		fromAmt, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			return "", err
		}
		fromAmtBf := new(big.Float).SetFloat64(fromAmt)
		res := fromAmtBf.Mul(fromAmtBf, big.NewFloat(NundPow))
		return res.Text('f', 9) + to, nil
	}

	return "", nil
}
