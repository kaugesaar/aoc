package solutions

import (
	"fmt"
	"sort"

	"github.com/kaugesaar/aoc/internal/solver"
)

// registry holds all registered solutions, organized by year and day.
var registry = make(map[int]map[int]solver.Solver)

// Register adds a solution for a specific year and day.
func Register(year, day int, s solver.Solver) {
	if registry[year] == nil {
		registry[year] = make(map[int]solver.Solver)
	}
	registry[year][day] = s
}

// Get retrieves a solution for a specific year and day.
func Get(year, day int) (solver.Solver, bool) {
	yearSolutions, ok := registry[year]
	if !ok {
		return nil, false
	}
	s, ok := yearSolutions[day]
	return s, ok
}

// GetAllForYear retrieves all solutions for a given year.
func GetAllForYear(year int) (map[int]solver.Solver, error) {
	yearSolutions, ok := registry[year]
	if !ok {
		return nil, fmt.Errorf("no solutions registered for year %d", year)
	}
	return yearSolutions, nil
}

// RegisteredYears returns a sorted list of years with registered solutions.
func RegisteredYears() []int {
	years := make([]int, 0, len(registry))
	for year := range registry {
		years = append(years, year)
	}
	sort.Ints(years)
	return years
}

// RegisteredDays returns a sorted list of days with registered solutions for a given year.
func RegisteredDays(year int) []int {
	yearSolutions, ok := registry[year]
	if !ok {
		return nil
	}
	days := make([]int, 0, len(yearSolutions))
	for day := range yearSolutions {
		days = append(days, day)
	}
	sort.Ints(days)
	return days
}
