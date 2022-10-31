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

func BenchmarkMist_Generate(b *testing.B) {
	b.Run("run_mist_gen", func(b *testing.B) {
		mist := NewMist()
		for i := 0; i < b.N; i++ {
			mist.Generate()
		}
	})
}
