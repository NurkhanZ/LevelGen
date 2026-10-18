package cli

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/NurkhanZ/LevelGen/internal/domain"
	"github.com/NurkhanZ/LevelGen/internal/generator"
	"github.com/NurkhanZ/LevelGen/internal/store"
)

// Run returns a process exit code and keeps all input and output injectable.
func Run(args []string, in io.Reader, out, errOut io.Writer, gen generator.Generator) int {
	fs := flag.NewFlagSet("level-gen", flag.ContinueOnError)
	fs.SetOutput(errOut)
	p := domain.LevelParams{}
	fs.IntVar(&p.Width, "width", 30, "level width (30-400)")
	fs.IntVar(&p.Height, "height", 30, "level height (30-400)")
	fs.Float64Var(&p.PlatformDensity, "platform-density", 0, "reserved platform density (0-1)")
	fs.Float64Var(&p.SpikeDensity, "spike-density", 0, "reserved spike density (0-1)")
	fs.IntVar(&p.Seed, "seed", 1, "generation seed")
	interactive := fs.Bool("interactive", false, "open the interactive menu")
	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return 0
		}
		return 2
	}
	if fs.NArg() != 0 {
		fmt.Fprintln(errOut, "unexpected positional arguments:", strings.Join(fs.Args(), " "))
		return 2
	}
	if err := p.Validate(); err != nil {
		fmt.Fprintln(errOut, err)
		return 1
	}
	if gen == nil {
		fmt.Fprintln(errOut, "generator is unavailable")
		return 1
	}
	if !*interactive {
		m, err := generate(gen, p)
		if err != nil {
			fmt.Fprintln(errOut, err)
			return 1
		}
		fmt.Fprint(out, m.String())
		return 0
	}
	return runInteractive(bufio.NewScanner(in), out, errOut, gen, p)
}

func generate(gen generator.Generator, p domain.LevelParams) (*domain.TileMap, error) {
	if err := p.Validate(); err != nil {
		return nil, err
	}
	m, err := gen.Generate(p)
	if err != nil {
		return nil, err
	}
	if m == nil {
		return nil, fmt.Errorf("generator returned no map")
	}
	return m, nil
}

func runInteractive(scan *bufio.Scanner, out, errOut io.Writer, gen generator.Generator, p domain.LevelParams) int {
	history := store.NewMemoryStore()
	fmt.Fprintln(out, "LevelGen dev — generate and inspect levels. History lasts for this run only.")
	for {
		fmt.Fprint(out, "\n1 Generate level\n2 Change settings\n3 View generated levels\n4 Quit\nChoice: ")
		choice, ok, err := readLine(scan)
		if err != nil {
			fmt.Fprintln(errOut, "input error:", err)
			return 1
		}
		if !ok {
			return 0
		}
		switch choice {
		case "1":
			m, err := generate(gen, p)
			if err != nil {
				fmt.Fprintln(out, "Generation error:", err)
				continue
			}
			task := domain.LevelTask{ID: strconv.Itoa(len(history.List()) + 1), Params: p, Status: domain.StatusPending}
			if err := task.Complete(m); err != nil {
				fmt.Fprintln(errOut, err)
				return 1
			}
			if err := history.Save(task); err != nil {
				fmt.Fprintln(errOut, err)
				return 1
			}
			preview, err := Preview(m, p)
			if err != nil {
				fmt.Fprintln(errOut, err)
				return 1
			}
			fmt.Fprintf(out, "Level %s\n%s", task.ID, preview)
		case "2":
			var eof bool
			p, eof, err = changeSettings(scan, out, p)
			if err != nil {
				fmt.Fprintln(errOut, "input error:", err)
				return 1
			}
			if eof {
				return 0
			}
		case "3":
			eof, err := viewHistory(scan, out, history)
			if err != nil {
				fmt.Fprintln(errOut, "input error:", err)
				return 1
			}
			if eof {
				return 0
			}
		case "4":
			return 0
		default:
			fmt.Fprintln(out, "Choose 1, 2, 3, or 4.")
		}
	}
}

func readLine(scan *bufio.Scanner) (string, bool, error) {
	if !scan.Scan() {
		return "", false, scan.Err()
	}
	return strings.TrimSpace(scan.Text()), true, nil
}

func changeSettings(scan *bufio.Scanner, out io.Writer, p domain.LevelParams) (domain.LevelParams, bool, error) {
	for {
		fmt.Fprintf(out, "Settings: width=%d height=%d platform-density=%g spike-density=%g seed=%d\n", p.Width, p.Height, p.PlatformDensity, p.SpikeDensity, p.Seed)
		fmt.Fprint(out, "1 Width  2 Height  3 Platform density  4 Spike density  5 Seed  6 Back\nSetting: ")
		choice, ok, err := readLine(scan)
		if err != nil || !ok {
			return p, !ok, err
		}
		if choice == "6" {
			return p, false, nil
		}
		if choice < "1" || choice > "5" || len(choice) != 1 {
			fmt.Fprintln(out, "Choose a setting from 1 to 6.")
			continue
		}
		fmt.Fprint(out, "New value: ")
		value, ok, err := readLine(scan)
		if err != nil || !ok {
			return p, !ok, err
		}
		next := p
		if choice == "3" || choice == "4" {
			n, err := strconv.ParseFloat(value, 64)
			if err != nil {
				fmt.Fprintln(out, "Enter a decimal number.")
				continue
			}
			if choice == "3" {
				next.PlatformDensity = n
			} else {
				next.SpikeDensity = n
			}
		} else {
			n, err := strconv.Atoi(value)
			if err != nil {
				fmt.Fprintln(out, "Enter a whole number.")
				continue
			}
			switch choice {
			case "1":
				next.Width = n
			case "2":
				next.Height = n
			case "5":
				next.Seed = n
			}
		}
		if err := next.Validate(); err != nil {
			fmt.Fprintln(out, err)
			continue
		}
		p = next
	}
}

func viewHistory(scan *bufio.Scanner, out io.Writer, history store.TaskStore) (bool, error) {
	tasks := history.List()
	if len(tasks) == 0 {
		fmt.Fprintln(out, "No generated levels in this run.")
		return false, nil
	}
	for _, task := range tasks {
		fmt.Fprintf(out, "%s: %dx%d seed=%d %s\n", task.ID, task.Params.Width, task.Params.Height, task.Params.Seed, task.Status)
	}
	fmt.Fprint(out, "Level ID (Enter to go back): ")
	id, ok, err := readLine(scan)
	if err != nil || !ok {
		return !ok, err
	}
	if id == "" {
		return false, nil
	}
	task, err := history.GetByID(id)
	if err != nil {
		fmt.Fprintln(out, "Unknown level ID.")
		return false, nil
	}
	preview, err := Preview(task.Result, task.Params)
	if err != nil {
		return false, err
	}
	fmt.Fprint(out, preview)
	return false, nil
}
