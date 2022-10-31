package snoflake

import (
	"fmt"
	"testing"
)

func TestSnowflake_Generate(t *testing.T) {
	snow, _ := NewSnowflake(1)
	for i := 0; i < 10; i++ {
		id := snow.Generate()
		fmt.Println(id)
	}
}

func BenchmarkNewSnowflake(b *testing.B) {
	b.Run("run_snowflake", func(b *testing.B) {
		snow, err := NewSnowflake(1)
		if err != nil {
			b.Errorf("new snowflake error %s", err)
			b.Fatal()
		}
		for i := 0; i < b.N; i++ {
			snow.Generate()
		}
	})
}
