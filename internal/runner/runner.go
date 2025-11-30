package runner

import (
	"fmt"
	"io"
	"time"

	"github.com/kaugesaar/aoc/internal/input"
	"github.com/kaugesaar/aoc/internal/solutions"
)

// RunDay executes both parts for a single day.
func RunDay(year, day int) error {
	solver, ok := solutions.Get(year, day)
	if !ok {
		return fmt.Errorf("no solution registered for year %d day %d", year, day)
	}

	fmt.Printf("Year %d, Day %d:\n", year, day)

	if err := runPart(solver.Part1, year, day, 1); err != nil {
		return err
	}

	if err := runPart(solver.Part2, year, day, 2); err != nil {
		return err
	}

	return nil
}

// RunYear executes all registered days for a year.
func RunYear(year int) error {
	yearSolutions, err := solutions.GetAllForYear(year)
	if err != nil {
		return err
	}

	days := solutions.RegisteredDays(year)
	if len(days) == 0 {
		return fmt.Errorf("no solutions registered for year %d", year)
	}

	fmt.Printf("Year %d:\n", year)

	totalDuration := time.Duration(0)

	for _, day := range days {
		solver := yearSolutions[day]
		fmt.Printf("  Day %d:\n", day)

		duration, err := runPartWithTiming(solver.Part1, year, day, 1, "    ")
		if err != nil {
			return err
		}
		totalDuration += duration

		duration, err = runPartWithTiming(solver.Part2, year, day, 2, "    ")
		if err != nil {
			return err
		}
		totalDuration += duration
	}

	fmt.Println("---")
	fmt.Printf("Total: %v\n", totalDuration)

	return nil
}

// RunPart executes a specific part (1 or 2) for a day.
func RunPart(year, day, part int) error {
	solver, ok := solutions.Get(year, day)
	if !ok {
		return fmt.Errorf("no solution registered for year %d day %d", year, day)
	}

	if part != 1 && part != 2 {
		return fmt.Errorf("invalid part: %d (must be 1 or 2)", part)
	}

	fmt.Printf("Year %d, Day %d, Part %d:\n", year, day, part)

	var fn func(io.Reader) (string, error)
	if part == 1 {
		fn = solver.Part1
	} else {
		fn = solver.Part2
	}

	return runPart(fn, year, day, part)
}

func runPart(fn func(io.Reader) (string, error), year, day, part int) error {
	_, err := runPartWithTiming(fn, year, day, part, "  ")
	return err
}

func runPartWithTiming(fn func(io.Reader) (string, error), year, day, part int, indent string) (time.Duration, error) {
	inputReader, err := input.GetInput(year, day)
	if err != nil {
		return 0, fmt.Errorf("failed to get input: %w", err)
	}

	start := time.Now()
	result, err := fn(inputReader)
	duration := time.Since(start)

	if err != nil {
		return 0, fmt.Errorf("part %d failed: %w", part, err)
	}

	fmt.Printf("%sPart %d: %s (%v)\n", indent, part, result, duration)
	return duration, nil
}
