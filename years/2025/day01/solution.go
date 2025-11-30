package day01

import (
	"bufio"
	"io"

	"github.com/kaugesaar/aoc/internal/solutions"
)

type Solution struct{}

func New() *Solution {
	return &Solution{}
}

func (s *Solution) Part1(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	// TODO: Implement Part 1
	_ = scanner

	return "", nil
}

func (s *Solution) Part2(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	// TODO: Implement Part 2
	_ = scanner

	return "", nil
}

func init() {
	solutions.Register(2025, 1, New())
}
