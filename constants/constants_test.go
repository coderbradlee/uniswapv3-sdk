package constants

import (
	"fmt"
	"testing"
)

func TestT(t *testing.T) {
	for _, p := range FeePairs {
		fmt.Println(p.Fee0, p.Fee1)
	}
}
