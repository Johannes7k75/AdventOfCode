package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var input_test string

func TestSolveP1(t *testing.T) {
	out := solve(input_test)
	want := 14

	if out != want {
		t.Errorf("solveP1() = %v, want = %v", out, want)
	}
}

func TestSolveP2(t *testing.T) {
	out := solve2(input_test)
	want := 34

	if out != want {
		t.Errorf("solveP2() = %v, want = %v", out, want)
	}
}
