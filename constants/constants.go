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
	result = Permutations(FeesInt, n)
	return
}

func Permutations(L []int, r int) [][]int {
	if r == 1 {
		//Convert every item in L to List and
		//Append it to List of List
		temp := make([][]int, 0)
		for _, rr := range L {
			t := make([]int, 0)
			t = append(t, rr)
			temp = append(temp, [][]int{t}...)
		}
		return temp
	} else {
		res := make([][]int, 0)
		for i := 0; i < len(L); i++ {
			//Create List Without L[i] element
			perms := make([]int, 0)
			perms = append(perms, L[:i]...)
			perms = append(perms, L[i+1:]...)
			//Call recursively to Permutations
			for _, x := range Permutations(perms, r-1) {
				t := append(x, L[i])
				res = append(res, [][]int{t}...)
			}
		}
		return res
	}
}
