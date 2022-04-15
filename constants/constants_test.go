package constants

import (
	"fmt"
	"testing"
)

func TestT(t *testing.T) {
	//for _, p := range FeePairs {
	//	fmt.Println(p.Fee0, p.Fee1)
	//}
	//result := make([][]int, 0)
	result, err := GetFeePermutations(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, r := range result {
		for _, tt := range r {
			fmt.Printf("%d ", tt)
		}
		fmt.Println("")
	}
}
