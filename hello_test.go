package testmod

import "testing"
import "fmt"

func TestHello(t *testing.T) {
	want := "Hello, world."
	fmt.Printf("test...\n")
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
