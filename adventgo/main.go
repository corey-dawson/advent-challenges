package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"adventgo.com/puzzles/day01"
	"adventgo.com/puzzles/internal/utils"
)

const (
	ROOT_DIR_NM = "advent-challenges"
)

// SolveFunc is the type signature for all solve functions
type SolveFunc func(string)

// DaySolver contains both solve functions for a day
type DaySolver struct {
	Solve1 SolveFunc
	Solve2 SolveFunc
}

// Registry maps day numbers (as strings like "01", "02") to their solvers
var dayRegistry = map[string]DaySolver{
	"01": {Solve1: day01.Solve1, Solve2: day01.Solve2},
	// Add more days here as you create them:
	// "02": {Solve1: day02.Solve1, Solve2: day02.Solve2},
	// "03": {Solve1: day03.Solve1, Solve2: day03.Solve2},
	// ... up to "25"
}

func main() {
	inputFile := flag.String("input", "sample", "Input file name (sample, input1, or input2)")
	problemDay := flag.String("day", "01", "Day number (01, 02, etc)")
	part := flag.Int("part", 1, "Part number (1 or 2)")
	flag.Parse()

	// Validate part number
	if *part != 1 && *part != 2 {
		fmt.Fprintf(os.Stderr, "Error: part must be 1 or 2, got %d\n", *part)
		os.Exit(1)
	}

	// Check if day exists in registry
	solver, exists := dayRegistry[*problemDay]
	if !exists {
		fmt.Fprintf(os.Stderr, "Error: day %s not found in registry. Available days: ", *problemDay)
		first := true
		for day := range dayRegistry {
			if !first {
				fmt.Fprintf(os.Stderr, ", ")
			}
			fmt.Fprintf(os.Stderr, "%s", day)
			first = false
		}
		fmt.Fprintf(os.Stderr, "\n")
		os.Exit(1)
	}

	// Get the appropriate solve function
	var solveFunc SolveFunc
	if *part == 1 {
		solveFunc = solver.Solve1
	} else {
		solveFunc = solver.Solve2
	}

	// Check if solve function is nil (not implemented yet)
	if solveFunc == nil {
		fmt.Fprintf(os.Stderr, "Error: day %s part %d is not implemented yet\n", *problemDay, *part)
		os.Exit(1)
	}

	// Setup paths to data
	dayFormat := fmt.Sprintf("day_%s", *problemDay)
	projRoot, err := utils.FindRootPath()
	if err != nil {
		panic(err)
	}
	inputPath := filepath.Join(projRoot, "data", dayFormat, *inputFile)

	// Display info
	utils.InitDisplay(dayFormat, *part, inputPath)

	// Call the appropriate solve function
	solveFunc(inputPath)
}
