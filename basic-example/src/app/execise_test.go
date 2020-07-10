package main

import "testing"

func TestAdd(t *testing.T) {
	var c int
	c = add(10, 5)
	if c != 15 {
		t.Error("Exepected 15,  got ", c)
	}
}

func TestSub(t *testing.T) {
	var c int
	c = sub(10, 5)
	if c != 5 {
		t.Error("Exepected 5,  got ", c)
	}
}

func TestMul(t *testing.T) {
	var c int
	c = mul(10, 5)
	if c != 50 {
		t.Error("Exepected 50,  got ", c)
	}
}

func TestDiv(t *testing.T) {
	var c int
	c = div(10, 5)
	if c != 2 {
		t.Error("Exepected 2,  got ", c)
	}
}
