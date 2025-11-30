package day01

import (
	"strings"
	"testing"
)

const sampleInput = `
TODO: Paste sample input here
`

func TestPart1(t *testing.T) {
	s := New()
	result, err := s.Part1(strings.NewReader(sampleInput))
	if err != nil {
		t.Fatalf("Part1 returned error: %v", err)
	}

	expected := "TODO: Put expected result here"
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

	expected := "TODO: Put expected result here"
	if result != expected {
		t.Errorf("Part2 = %v, want %v", result, expected)
	}
}
