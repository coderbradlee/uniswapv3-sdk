package constants

import (
	"fmt"
	"testing"
)

func TestT(t *testing.T) {
	//for _, p := range FeePairs {
	//	fmt.Println(p.Fee0, p.Fee1)
	//}
	result := make([][]FeeAmount, 0)
	generatePermutations(Fees, 3, &result)
	for _, r := range result {
		for _, tt := range r {
			fmt.Printf("%d ", tt)
		}
		fmt.Println("")
	}
}
