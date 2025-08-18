package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// options - struct with grep flags
type options struct {
	after     int
	before    int
	countOnly bool
	ignore    bool
	invert    bool
	fixed     bool
	lineNum   bool
	pattern   string
	file      string
}

// line - struct that contains num of line and it's containing
type line struct {
	num  int
	text string
}

func main() {
	opts := parseFlags()

	var input io.Reader = os.Stdin
	if opts.file != "" {
		f, err := os.Open(opts.file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to open file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		input = f
	}

	runGrep(input, opts)
}

// parseFlags - read flags and returns options
func parseFlags() options {
	A := flag.Int("A", 0, "print N lines after match")
	B := flag.Int("B", 0, "print N lines before match")
	C := flag.Int("C", 0, "print N lines around match")
	c := flag.Bool("c", false, "print only count of matching lines")
	i := flag.Bool("i", false, "ignore case")
	v := flag.Bool("v", false, "invert match")
	F := flag.Bool("F", false, "treat pattern as fixed string")
	n := flag.Bool("n", false, "print line numbers")

	flag.Parse()

	// override A and B
	if *C > 0 {
		*A, *B = *C, *C
	}

	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Usage: grep [flags] PATTERN [FILE]")
		os.Exit(1)
	}

	pattern := args[0]
	file := ""
	if len(args) > 1 {
		file = args[1]
	}

	return options{
		after:     *A,
		before:    *B,
		countOnly: *c,
		ignore:    *i,
		invert:    *v,
		fixed:     *F,
		lineNum:   *n,
		pattern:   pattern,
		file:      file,
	}
}

func runGrep(r io.Reader, opts options) {
	scanner := bufio.NewScanner(r)

	var re *regexp.Regexp
	if !opts.fixed {
		flags := ""
		if opts.ignore {
			flags = "(?i)"
		}
		var err error
		re, err = regexp.Compile(flags + opts.pattern)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid regex: %v\n", err)
			os.Exit(1)
		}
	} else if opts.ignore {
		opts.pattern = strings.ToLower(opts.pattern)
	}

	var beforeBuf []line
	afterRemain := 0
	printed := make(map[int]bool)
	count := 0

	lineNum := 0
	for scanner.Scan() {
		lineNum++
		text := scanner.Text()

		// check match
		match := false
		if opts.fixed {
			src := text
			if opts.ignore {
				src = strings.ToLower(src)
			}
			match = strings.Contains(src, opts.pattern)
		} else {
			match = re.MatchString(text)
		}

		// invert match
		if opts.invert {
			match = !match
		}


		if match {
			count++
			if opts.countOnly {
				afterRemain = opts.after
				beforeBuf = nil
				continue
			}

			// print before context
			for _, l := range beforeBuf {
				if !printed[l.num] {
					printLine(l, opts)
					printed[l.num] = true
				}
			}
			beforeBuf = nil

			// print current line
			if !printed[lineNum] {
				printLine(line{num: lineNum, text: text}, opts)
				printed[lineNum] = true
			}
			afterRemain = opts.after
		} else {
			// after context
			if afterRemain > 0 {
				if !printed[lineNum] {
					printLine(line{num: lineNum, text: text}, opts)
					printed[lineNum] = true
				}
				afterRemain--
			}

			// before buffer
			if opts.before > 0 {
				if len(beforeBuf) >= opts.before {
					beforeBuf = beforeBuf[1:]
				}
				beforeBuf = append(beforeBuf, line{num: lineNum, text: text})
			}
		}
	}

	if opts.countOnly {
		fmt.Println(count)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "read error: %v\n", err)
		os.Exit(1)
	}
}


var out io.Writer = os.Stdout

func printLine(l line, opts options) {
	if opts.lineNum {
		fmt.Fprintf(out, "%d:%s\n", l.num, l.text)
	} else {
		fmt.Fprintln(out, l.text)
	}
}
