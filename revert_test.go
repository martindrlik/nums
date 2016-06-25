package main

import (
	"testing"
)

func TestRevert(t *testing.T) {
	s := Revert("hello world")
	if s != "dlrow olleh" {
		t.Errorf("%q should be %q", s, "dlrow olleh")
	}
}
