package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Hezza")
	want := "Hello, Hezza"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
