package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Raul")
	want := "Hello, Raul!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
