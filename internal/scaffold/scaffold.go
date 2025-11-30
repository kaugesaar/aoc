package scaffold

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kaugesaar/aoc/internal/auth"
	"github.com/kaugesaar/aoc/internal/input"
)

// Create generates boilerplate for a new day.
func Create(year, day int) error {
	dayDir := filepath.Join("years", fmt.Sprintf("%d", year), fmt.Sprintf("day%02d", day))
	if err := os.MkdirAll(dayDir, 0o755); err != nil {
		return fmt.Errorf("failed to create day directory: %w", err)
	}

	solutionPath := filepath.Join(dayDir, "solution.go")
	if _, err := os.Stat(solutionPath); err == nil {
		return fmt.Errorf("solution already exists: %s", solutionPath)
	}

	template := generateSolutionTemplate(year, day)
	if err := os.WriteFile(solutionPath, []byte(template), 0o644); err != nil {
		return fmt.Errorf("failed to write solution.go: %w", err)
	}

	testPath := filepath.Join(dayDir, "solution_test.go")
	testTemplate := generateTestTemplate(day)
	if err := os.WriteFile(testPath, []byte(testTemplate), 0o644); err != nil {
		return fmt.Errorf("failed to write solution_test.go: %w", err)
	}

	if auth.IsConfigured() {
		if err := input.Download(year, day); err != nil {
			fmt.Printf("Warning: failed to download input: %v\n", err)
		} else {
			fmt.Printf("Downloaded input for %d day %d\n", year, day)
		}
	} else {
		fmt.Println("Session not configured - run 'aoc login' to enable automatic input downloads")
	}

	if err := updateYearFile(year, day); err != nil {
		return fmt.Errorf("failed to update year.go: %w", err)
	}

	fmt.Printf("Created solution at %s\n", solutionPath)
	fmt.Printf("Created test at %s\n", testPath)
	return nil
}

func generateSolutionTemplate(year, day int) string {
	return fmt.Sprintf(`package day%02d

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
	solutions.Register(%d, %d, New())
}
`, day, year, day)
}

func generateTestTemplate(day int) string {
	return fmt.Sprintf(`package day%02d

import (
	"strings"
	"testing"
)

const sampleInput = `+"`"+`
TODO: Paste sample input here
`+"`"+`

func TestPart1(t *testing.T) {
	s := New()
	result, err := s.Part1(strings.NewReader(sampleInput))
	if err != nil {
		t.Fatalf("Part1 returned error: %%v", err)
	}

	expected := "TODO: Put expected result here"
	if result != expected {
		t.Errorf("Part1 = %%v, want %%v", result, expected)
	}
}

func TestPart2(t *testing.T) {
	s := New()
	result, err := s.Part2(strings.NewReader(sampleInput))
	if err != nil {
		t.Fatalf("Part2 returned error: %%v", err)
	}

	expected := "TODO: Put expected result here"
	if result != expected {
		t.Errorf("Part2 = %%v, want %%v", result, expected)
	}
}
`, day)
}

func updateYearFile(year, day int) error {
	yearDir := filepath.Join("years", fmt.Sprintf("%d", year))
	yearFile := filepath.Join(yearDir, "year.go")

	if err := os.MkdirAll(yearDir, 0o755); err != nil {
		return err
	}

	importLine := fmt.Sprintf("\t_ \"github.com/kaugesaar/aoc/years/%d/day%02d\"", year, day)

	content, err := os.ReadFile(yearFile)
	if os.IsNotExist(err) {
		newContent := fmt.Sprintf(`package year%d

import (
%s
)
`, year, importLine)
		return os.WriteFile(yearFile, []byte(newContent), 0o644)
	}

	if err != nil {
		return err
	}

	if strings.Contains(string(content), importLine) {
		return nil
	}

	lines := strings.Split(string(content), "\n")
	var imports []string
	var beforeImport []string
	var afterImport []string
	inImportBlock := false
	importBlockFound := false

	for _, line := range lines {
		if !importBlockFound && strings.Contains(line, "import (") {
			inImportBlock = true
			importBlockFound = true
			beforeImport = append(beforeImport, line)
			continue
		}
		if inImportBlock && strings.Contains(line, ")") {
			inImportBlock = false
			continue
		}
		if inImportBlock {
			trimmed := strings.TrimSpace(line)
			if trimmed != "" {
				imports = append(imports, trimmed)
			}
		} else if !importBlockFound {
			beforeImport = append(beforeImport, line)
		} else {
			afterImport = append(afterImport, line)
		}
	}

	newImport := strings.TrimSpace(importLine)
	imports = append(imports, newImport)
	sortImports(imports)

	var result []string
	result = append(result, beforeImport...)
	for _, imp := range imports {
		result = append(result, "\t"+imp)
	}
	result = append(result, ")")
	result = append(result, afterImport...)

	return os.WriteFile(yearFile, []byte(strings.Join(result, "\n")), 0o644)
}

func sortImports(imports []string) {
	for i := 0; i < len(imports); i++ {
		for j := i + 1; j < len(imports); j++ {
			if imports[i] > imports[j] {
				imports[i], imports[j] = imports[j], imports[i]
			}
		}
	}
}
