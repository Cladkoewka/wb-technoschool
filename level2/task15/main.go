package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// redirect SIGINT to shell
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT)

	fmt.Println("minishell. Ctrl+D — exit, Ctrl+C — interrupt command.")
	for {
		wd, _ := os.Getwd()
		fmt.Printf("%s$ ", wd)

		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("\nexit")
				return
			}
			fmt.Fprintf(os.Stderr, "read error: %v\n", err)
			continue
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// split by pipe, '|' divide commands
		parts := splitByUnescapedPipe(line)
		var pipeline [][]string
		parseErr := false
		for _, p := range parts {
			fields := simpleFields(strings.TrimSpace(p))
			if len(fields) == 0 {
				fmt.Fprintln(os.Stderr, "syntax error: empty command in pipeline")
				parseErr = true
				break
			}
			for i, a := range fields {
				fields[i] = substituteEnv(a)
			}
			pipeline = append(pipeline, fields)
		}
		if parseErr {
			continue
		}

		// execute pipline
		if len(pipeline) == 0 {
			continue
		}

		// for  correct redirect SIGINT build pids outer processes
		var pidsMu sync.Mutex
		var childPids []int

		// forward SIGINT to daughter pgroups (по pid'ам)
		stopForward := make(chan struct{})
		go func() {
			for {
				select {
				case <-sigc:
					pidsMu.Lock()
					for _, pid := range childPids {
						_ = syscall.Kill(-pid, syscall.SIGINT)
					}
					pidsMu.Unlock()
				case <-stopForward:
					return
				}
			}
		}()

		exitCode := runPipeline(pipeline, &pidsMu, &childPids)

		close(stopForward)

		_ = exitCode
	}
}

// splitByUnescapedPipe — divide by '|' (without escape/quotes)
func splitByUnescapedPipe(s string) []string {
	return strings.Split(s, "|")
}

// simpleFields — divide by space
func simpleFields(s string) []string {
	return strings.Fields(s)
}

// substituteEnv: replace $VAR and ${VAR} from env
func substituteEnv(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch != '$' {
			b.WriteByte(ch)
			continue
		}

		if i+1 < len(s) && s[i+1] == '{' {
			// search }
			j := i + 2
			for j < len(s) && s[j] != '}' {
				j++
			}
			if j < len(s) && s[j] == '}' {
				name := s[i+2 : j]
				b.WriteString(os.Getenv(name))
				i = j
				continue
			} else {
				b.WriteByte('$')
				continue
			}
		}
		// $VAR style
		j := i + 1
		if j < len(s) && isVarChar(s[j]) {
			for j < len(s) && isVarChar(s[j]) {
				j++
			}
			name := s[i+1 : j]
			b.WriteString(os.Getenv(name))
			i = j - 1
			continue
		}
		b.WriteByte('$')
	}
	return b.String()
}

func isVarChar(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9') || c == '_'
}

// runPipeline: pipeline — slice of commands (every command — []string, first - name).
// pidsMu и childPids for outer processes
func runPipeline(pipeline [][]string, pidsMu *sync.Mutex, childPids *[]int) int {
	n := len(pipeline)
	// if single command
	if n == 1 && isBuiltin(pipeline[0][0]) {
		cmd := buildCommand(pipeline[0])
		return runBuiltin(cmd, os.Stdin, os.Stdout)
	}

	type proc struct {
		name string
		args []string
		cmd  *exec.Cmd
	}
	procs := make([]*proc, n)
	for i := 0; i < n; i++ {
		procs[i] = &proc{
			name: pipeline[i][0],
			args: pipeline[i][1:],
		}
	}

	// pipes: reader/writer pairs
	type rw struct {
		r *os.File
		w *os.File
	}
	pipes := make([]rw, 0, n-1)
	for i := 0; i < n-1; i++ {
		r, w, err := os.Pipe()
		if err != nil {
			fmt.Fprintf(os.Stderr, "pipe error: %v\n", err)
			return 1
		}
		pipes = append(pipes, rw{r: r, w: w})
	}

	// Run every process
	var wg sync.WaitGroup
	exitCodes := make([]int, n)

	for i := 0; i < n; i++ {
		// stdin
		var stdin io.Reader = os.Stdin
		if i > 0 {
			stdin = pipes[i-1].r
		}
		// stdout
		var stdout io.Writer = os.Stdout
		if i < n-1 {
			stdout = pipes[i].w
		}

		name := procs[i].name
		args := procs[i].args

		if isBuiltin(name) {
			// run builtin ig goroutine
			wg.Add(1)
			go func(idx int, name string, args []string, in io.Reader, out io.Writer) {
				defer wg.Done()
				cmd := buildCommand(append([]string{name}, args...))
				code := runBuiltin(cmd, in, out)
				exitCodes[idx] = code
				if w, ok := out.(*os.File); ok {
					_ = w.Close()
				} else if pw, ok := out.(*os.File); ok {
					_ = pw.Close()
				}
			}(i, name, args, stdin, stdout)
			continue
		}

		// outer process
		wg.Add(1)
		go func(idx int, name string, args []string, in io.Reader, out io.Writer) {
			defer wg.Done()
			cmd := exec.Command(name, args...)
			cmd.Stdin = in
			cmd.Stdout = out
			cmd.Stderr = os.Stderr
			cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

			if err := cmd.Start(); err != nil {
				fmt.Fprintf(os.Stderr, "failed to start %s: %v\n", name, err)
				exitCodes[idx] = 127
				if w, ok := out.(*os.File); ok {
					_ = w.Close()
				}
				return
			}

			pidsMu.Lock()
			*childPids = append(*childPids, cmd.Process.Pid)
			pidsMu.Unlock()

			err := cmd.Wait()
			if err != nil {
				if exitErr, ok := err.(*exec.ExitError); ok {
					if ws, ok := exitErr.Sys().(syscall.WaitStatus); ok {
						exitCodes[idx] = ws.ExitStatus()
					} else {
						exitCodes[idx] = 1
					}
				} else {
					exitCodes[idx] = 1
				}
			} else {
				exitCodes[idx] = 0
			}
			if w, ok := out.(*os.File); ok {
				_ = w.Close()
			}
		}(i, name, args, stdin, stdout)
	}

	// close all writers
	for _, p := range pipes {
		_ = p.w.Close()
	}

	wg.Wait()

	// close all read ends
	for _, p := range pipes {
		_ = p.r.Close()
	}

	if len(exitCodes) == 0 {
		return 0
	}
	return exitCodes[len(exitCodes)-1]
}

// Command - helper for builtin
type Command struct {
	Name string
	Args []string
}

func buildCommand(fields []string) *Command {
	return &Command{Name: fields[0], Args: fields[1:]}
}

func isBuiltin(name string) bool {
	switch name {
	case "cd", "pwd", "echo", "kill", "ps":
		return true
	default:
		return false
	}
}

// runBuiltin: execute builtin command
func runBuiltin(cmd *Command, in io.Reader, out io.Writer) int {
	switch cmd.Name {
	case "cd":
		path := "."
		if len(cmd.Args) > 0 {
			path = cmd.Args[0]
		}
		if err := os.Chdir(path); err != nil {
			fmt.Fprintf(out, "cd: %v\n", err)
			return 1
		}
		return 0
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(out, "pwd: %v\n", err)
			return 1
		}
		fmt.Fprintln(out, dir)
		return 0
	case "echo":
		for i, a := range cmd.Args {
			if i > 0 {
				fmt.Fprint(out, " ")
			}
			fmt.Fprint(out, a)
		}
		fmt.Fprintln(out)
		return 0
	case "kill":
		if len(cmd.Args) < 1 {
			fmt.Fprintln(out, "kill: expected pid")
			return 1
		}
		pidStr := cmd.Args[0]
		pid, err := strconv.Atoi(pidStr)
		if err != nil {
			fmt.Fprintf(out, "kill: invalid pid: %v\n", err)
			return 1
		}
		if err := syscall.Kill(pid, syscall.SIGTERM); err != nil {
			fmt.Fprintf(out, "kill: %v\n", err)
			return 1
		}
		return 0
	case "ps":
		c := exec.Command("ps", "-eo", "pid,cmd")
		c.Stdout = out
		c.Stderr = os.Stderr
		_ = c.Run()
		return 0
	default:
		fmt.Fprintf(out, "unknown builtin: %s\n", cmd.Name)
		return 1
	}
}
