package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/kaugesaar/aoc/internal/auth"
	"github.com/kaugesaar/aoc/internal/input"
	"github.com/kaugesaar/aoc/internal/runner"
	"github.com/kaugesaar/aoc/internal/scaffold"

	_ "github.com/kaugesaar/aoc/years/2025"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "login":
		handleLogin()
	case "scaffold":
		handleScaffold()
	case "run":
		handleRun()
	case "input":
		handleInput()
	default:
		fmt.Printf("Unknown command: %s\n\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: aoc <command> [options]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  login                              Save session cookie")
	fmt.Println("  scaffold -year <year> -day <day>   Generate boilerplate for a new day")
	fmt.Println("  run -year <year> [-day <day>] [-part <part>]")
	fmt.Println("                                     Run solutions")
	fmt.Println("  input -year <year> -day <day>      Download input")
}

func handleLogin() {
	fmt.Print("Enter your Advent of Code session cookie: ")
	reader := bufio.NewReader(os.Stdin)
	cookie, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	cookie = strings.TrimSpace(cookie)
	if cookie == "" {
		fmt.Println("Error: session cookie cannot be empty")
		os.Exit(1)
	}

	if err := auth.SaveSession(cookie); err != nil {
		fmt.Printf("Error saving session: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Session saved successfully!")
}

func handleScaffold() {
	fs := flag.NewFlagSet("scaffold", flag.ExitOnError)
	year := fs.Int("year", 0, "Year")
	day := fs.Int("day", 0, "Day")
	fs.Parse(os.Args[2:])

	if *year == 0 || *day == 0 {
		fmt.Println("Error: -year and -day are required")
		fs.Usage()
		os.Exit(1)
	}

	if err := scaffold.Create(*year, *day); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func handleRun() {
	fs := flag.NewFlagSet("run", flag.ExitOnError)
	year := fs.Int("year", 0, "Year")
	day := fs.Int("day", 0, "Day (optional - runs all days if not specified)")
	part := fs.Int("part", 0, "Part (1 or 2, optional)")
	fs.Parse(os.Args[2:])

	if *year == 0 {
		fmt.Println("Error: -year is required")
		fs.Usage()
		os.Exit(1)
	}

	var err error

	if *day == 0 {
		err = runner.RunYear(*year)
	} else if *part == 0 {
		err = runner.RunDay(*year, *day)
	} else {
		err = runner.RunPart(*year, *day, *part)
	}

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func handleInput() {
	fs := flag.NewFlagSet("input", flag.ExitOnError)
	year := fs.Int("year", 0, "Year")
	day := fs.Int("day", 0, "Day")
	fs.Parse(os.Args[2:])

	if *year == 0 || *day == 0 {
		fmt.Println("Error: -year and -day are required")
		fs.Usage()
		os.Exit(1)
	}

	if err := input.Download(*year, *day); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Downloaded input for %d day %d\n", *year, *day)
}
