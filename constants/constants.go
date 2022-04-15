package constants

import (
	"errors"
	"math/big"

	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/common"
)

const PoolInitCodeHash = "0xe34f199b19b2b4f47f68442619d555527d244f78a3297ea89325f843f87b8b54"

var (
	FactoryAddress = common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984")
	AddressZero    = common.HexToAddress("0x0000000000000000000000000000000000000000")
)

// The default factory enabled fee amounts, denominated in hundredths of bips.
type FeeAmount uint64

const (
	FeeLowest FeeAmount = 100
	FeeLow    FeeAmount = 500
	FeeMedium FeeAmount = 3000
	FeeHigh   FeeAmount = 10000

	FeeMax FeeAmount = 1000000
)

type FeePair struct {
	Fee0 FeeAmount
	Fee1 FeeAmount
}

var (
	Fees     = []FeeAmount{FeeLow, FeeMedium, FeeHigh}
	FeesInt  = []int{int(FeeLow), int(FeeMedium), int(FeeHigh)}
	FeePairs = make([]FeePair, 0)
)

func init() {
	for i := 0; i < len(Fees); i++ {
		for j := 0; j < len(Fees); j++ {
			if i != j {
				FeePairs = append(FeePairs, FeePair{
					Fee0: Fees[i],
					Fee1: Fees[j],
				})
			}
		}
	}
}

// The default factory tick spacings by fee amount.
var TickSpacings = map[FeeAmount]int{
	FeeLowest: 1,
	FeeLow:    10,
	FeeMedium: 60,
	FeeHigh:   200,
}

var (
	NegativeOne = big.NewInt(-1)
	Zero        = big.NewInt(0)
	One         = big.NewInt(1)

	// used in liquidity amount math
	Q96  = new(big.Int).Exp(big.NewInt(2), big.NewInt(96), nil)
	Q192 = new(big.Int).Exp(Q96, big.NewInt(2), nil)

	PercentZero = entities.NewFraction(big.NewInt(0), big.NewInt(1))
)

func GetFeePermutations(n int) (result [][]int, err error) {
	if n == 0 || n > 3 {
		err = errors.New("n invalid")
		return
	}
	result = make([][]int, 0)
	generatePermutations(FeesInt, n, &result)
	return
}

func generatePermutations(array []int, n int, result *[][]int) {
	if n == 1 {
		dst := make([]int, len(array))
		copy(dst, array[:])
		*result = append(*result, dst)
	} else {
		for i := 0; i < n; i++ {
			generatePermutations(array, n-1, result)
			if n%2 == 0 {
				// Golang allow us to do multiple assignments
				array[0], array[n-1] = array[n-1], array[0]
			} else {
				array[i], array[n-1] = array[n-1], array[i]
			}
		}
	}
}
