package main

import "testing"
import "image/color"

func TestOrigin(t *testing.T) {
	p := cPoint{color.White, 0, 0}
	c := p.toComplex(-1, 1, -1, 1)

	if real(c) != -1 || imag(c) != -1 {
		t.Error("Wrong conversion", c, " expected (-1,-1i)")
	}
}

func TestCenter(t *testing.T) {
	p := cPoint{color.White, width / 2, height / 2}
	c := p.toComplex(-1, 1, -1, 1)

	if real(c) != 0 || imag(c) != 0 {
		t.Error("Wrong conversion", c, " expected (-1,-1i)")
	}
}

func TestCorner(t *testing.T) {
	p := cPoint{color.White, width, height}
	c := p.toComplex(-1, 1, -1, 1)

	if real(c) != 1 || imag(c) != 1 {
		t.Error("Wrong conversion", c, " expected (-1,-1i)")
	}
}
