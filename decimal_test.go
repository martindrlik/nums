package main

import (
	"testing"
)

func TestParseDecimal(t *testing.T) {
	d := ParseDecimal("10F", 16)
	if d != 271 {
		t.Errorf("(10F)16 should be 271 not %d", d)
	}
}

func TestConv(t *testing.T) {
	d := Decimal(7)
	s := d.Conv(2)
	if s != "111" {
		t.Errorf("7 should be (111)2 not (%s)2", s)
	}
}

func TestConv0(t *testing.T) {
	d := Decimal(0)
	s := d.Conv(8)
	if s != "0" {
		t.Errorf("0 should be (0)8 not (%s)8", s)
	}
}
