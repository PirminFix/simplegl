package main

import "testing"

func TestSizeofFloat32(t *testing.T) {

	if sizeof([]float32{1, 2, 3, 4}) != 4*4 {
		t.Fail()
	}
}
