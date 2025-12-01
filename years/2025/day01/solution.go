package day01

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kaugesaar/aoc/internal/solutions"
	"github.com/kaugesaar/aoc/pkg/strutil"
)

const (
	startPos  = 50
	dialRange = 100
)

type Solution struct{}

func New() *Solution {
	return &Solution{}
}

func (s *Solution) Part1(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	count, pos := 0, startPos

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		dir, steps := line[0], strutil.ToInt(line[1:])

		if dir == 'L' {
			pos -= steps
		} else {
			pos += steps
		}

		pos = ((pos % dialRange) + dialRange) % dialRange
		if pos == 0 {
			count++
		}
	}

	return fmt.Sprintf("%d", count), nil
}

func (s *Solution) Part2(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	count, pos := 0, startPos

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		dir, steps := line[0], strutil.ToInt(line[1:])

		count += steps / dialRange

		remainder := steps % dialRange

		if pos != 0 {
			if dir == 'L' && remainder >= pos {
				count++
			} else if dir == 'R' && pos+remainder >= dialRange {
				count++
			}
		}

		if dir == 'L' {
			pos -= steps
		} else {
			pos += steps
		}

		pos = ((pos % dialRange) + dialRange) % dialRange
	}

	return fmt.Sprintf("%d", count), nil
}

func init() {
	solutions.Register(2025, 1, New())
}
