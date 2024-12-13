package basics

import "testing"

func TestHello(t *testing.T) {
	got := HelloWorld()
	want := "Hello, World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
