package mist

import (
	"fmt"
	"testing"
)

func TestMist_Generate(t *testing.T) {
	mist := NewMist()
	for i := 0; i < 100; i++ {
		fmt.Println(mist.Generate())
	}
}
