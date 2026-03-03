package main

import "testing"

func TestGreeting(t *testing.T) {
	want := "Hello World!"
	got := greet()
	print(want, got)
}
