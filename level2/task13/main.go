package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type config struct {
	fields    map[int]struct{}
	delimiter string
	separated bool
}

// parseFields - divide string like "1,3-5,7"
func parseFields(spec string) (map[int]struct{}, error) {
	fields := make(map[int]struct{})

	parts := strings.Split(spec, ",")
	for _, p := range parts {
		if strings.Contains(p, "-") {
			bounds := strings.SplitN(p, "-", 2)
			if len(bounds) != 2 {
				return nil, fmt.Errorf("invalid range: %s", p)
			}
			start, err1 := strconv.Atoi(bounds[0])
			end, err2 := strconv.Atoi(bounds[1])
			if err1 != nil || err2 != nil || start <= 0 || end < start {
				return nil, fmt.Errorf("invalid range: %s", p)
			}
			for i := start; i <= end; i++ {
				fields[i] = struct{}{}
			}
		} else {
			num, err := strconv.Atoi(p)
			if err != nil || num <= 0 {
				return nil, fmt.Errorf("invalid field: %s", p)
			}
			fields[num] = struct{}{}
		}
	}

	return fields, nil
}

// parseFlags - returns config
func parseFlags() (*config, error) {
	fFlag := flag.String("f", "", "fields (like: 1,3-5)")
	dFlag := flag.String("d", "\t", "delimiter (defaul \\t)")
	sFlag := flag.Bool("s", false, "if true, output only lines with divider")
	flag.Parse()

	if *fFlag == "" {
		return nil, fmt.Errorf("-f needed")
	}

	fields, err := parseFields(*fFlag)
	if err != nil {
		return nil, err
	}

	return &config{
		fields:    fields,
		delimiter: *dFlag,
		separated: *sFlag,
	}, nil
}

// processLine - process one line
func processLine(line string, cfg *config) (string, bool) {
	if cfg.separated && !strings.Contains(line, cfg.delimiter) {
		return "", false
	}

	cols := strings.Split(line, cfg.delimiter)
	var out []string

	for i := 1; i <= len(cols); i++ {
		if _, ok := cfg.fields[i]; ok {
			out = append(out, cols[i-1])
		}
	}

	if len(out) == 0 {
		return "", false
	}
	return strings.Join(out, cfg.delimiter), true
}

// run starts processing STDIN
func run(cfg *config) error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if res, ok := processLine(line, cfg); ok {
			fmt.Println(res)
		}
	}
	return scanner.Err()
}

func main() {
	cfg, err := parseFlags()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	if err := run(cfg); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
