package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// main reads numeric data from a file and prints Linear Regression Line and Pearson Correlation Coefficient
func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run your-program.go <filename>")
		os.Exit(1)
	}

	yVals, err := readData(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	n := float64(len(yVals))
	var sumX, sumY, sumXY, sumX2, sumY2 float64

	for i, y := range yVals {
		x := float64(i)
		sumX += x
		sumY += y
		sumXY += x * y
		sumX2 += x * x
		sumY2 += y * y
	}

	// Linear regression: y = mx + b
	denom := n*sumX2 - sumX*sumX
	m := (n*sumXY - sumX*sumY) / denom
	b := (sumY*sumX2 - sumX*sumXY) / denom

	// Pearson correlation coefficient
	r := (n*sumXY - sumX*sumY) / math.Sqrt((n*sumX2-sumX*sumX)*(n*sumY2-sumY*sumY))

	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", m, b)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)
}

// readData reads a file and returns a slice of float64 values, one per line
func readData(path string) ([]float64, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var vals []float64
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		v, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value: %s", line)
		}
		vals = append(vals, v)
	}
	return vals, scanner.Err()
}
