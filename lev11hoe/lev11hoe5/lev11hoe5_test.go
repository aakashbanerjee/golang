package lev11hoe5

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	a := 10
	b := 20
	total := sum(a, b)
	fmt.Println(total)
	if total != 30 {
		t.Error("Expected 20, got : ", total)
	}
}
