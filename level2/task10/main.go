package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)


func main() {
	cfg := parseFlags()

	lines, err := readInput(cfg.files)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	if cfg.unique {
		lines = uniqueStrings(lines)
	}

	sort.Slice(lines, func(i, j int) bool {
		a := getColumn(lines[i], cfg.column)
		b := getColumn(lines[j], cfg.column)

		if cfg.numeric {
			af, _ := strconv.ParseFloat(a, 64)
			bf, _ := strconv.ParseFloat(b, 64)
			if cfg.reverse {
				return af > bf
			}
			return af < bf
		}

		if cfg.reverse {
			return a > b
		}
		return a < b
	})

	for _, line := range lines {
		fmt.Println(line)
	}
}

// Config contains flags
type Config struct {
	column  int  // -k: column number
	numeric bool // -n: 
	reverse bool // -r: 
	unique  bool // -u: 
	files   []string // input files
}

// parseFlags parse CLI args
func parseFlags() Config {
	cfg := Config{}
	flag.IntVar(&cfg.column, "k", 1, "Column number (>=1, divider — tab)")
	flag.BoolVar(&cfg.numeric, "n", false, "Numeric sort")
	flag.BoolVar(&cfg.reverse, "r", false, "Reverse sort")
	flag.BoolVar(&cfg.unique, "u", false, "Unique strings")
	flag.Parse()
	cfg.files = flag.Args()
	return cfg
}


// readInput читает данные из файла или stdin.
func readInput(files []string) ([]string, error) {
	var scanner *bufio.Scanner

	if len(files) > 0 {
		file, err := os.Open(files[0])
		if err != nil {
			return nil, err
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// getColumn returns N-th column (1-based)
func getColumn(line string, col int) string {
	parts := strings.Split(line, "\t")
	if col >= 1 && col <= len(parts) {
		return parts[col-1]
	}
	return ""
}

// uniqueStrings remove duplicates.
func uniqueStrings(input []string) []string {
	seen := make(map[string]struct{})
	var result []string
	for _, v := range input {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}