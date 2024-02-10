package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	fmt.Println("hello")
	t.Run("subtest", func(t *testing.T) {
		fmt.Println("subtest")
	})
}
