package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Adding 1 + 1 is 2", func(t *testing.T) {
		got := Add(1, 1)
		want := 2
		assertCorrectMessage(t, got, want)
	})

}
func assertCorrectMessage(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
}
