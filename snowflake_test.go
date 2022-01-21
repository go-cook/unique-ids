package mist

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
