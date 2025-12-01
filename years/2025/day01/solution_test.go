package day01

import (
	"strings"
	"testing"
)

const sampleInput = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

func TestPart1(t *testing.T) {
	s := New()
	result, err := s.Part1(strings.NewReader(sampleInput))
	if err != nil {
		t.Fatalf("Part1 returned error: %v", err)
	}

	expected := "3"
	if result != expected {
		t.Errorf("Part1 = %v, want %v", result, expected)
	}
}

func TestPart2(t *testing.T) {
	s := New()
	result, err := s.Part2(strings.NewReader(sampleInput))
	if err != nil {
		t.Fatalf("Part2 returned error: %v", err)
	}

	expected := "6"
	if result != expected {
		t.Errorf("Part2 = %v, want %v", result, expected)
	}
}
