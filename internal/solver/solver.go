package solver

import "io"

// Solver defines the interface that all daily solutions must implement.
type Solver interface {
	Part1(input io.Reader) (string, error)
	Part2(input io.Reader) (string, error)
}
