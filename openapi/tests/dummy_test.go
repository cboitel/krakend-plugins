package tests

import (
	"testing"
)

// TestDummy : does nothing but with success
func TestDummy(t *testing.T) {
	t.Log("Dummy test succeeded")
}

// TestAnotherDummy : another tests that does not fail
func TestAnotherDummy(t *testing.T) {
	t.Log("Yet another test suceeded")
	// useless code kept for example
	got := 1
	if got != 1 {
		t.Errorf("got = %d; want 1", got)
	}
}

func init() {
}
