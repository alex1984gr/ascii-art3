package main

import (
	// parse command-line flags
	"flag"
	// formatted I/O for errors
	"fmt"
	// io.Reader/Writer utilities
	"io"
	// OS-level helpers (stdin/stdout, exit)
	"os"
	// string helpers (Join, TrimRight)
	"strings"

	// local package that provides the ascii-art pipeline
	"ascii-art/pipeline"
)

// main is the program entrypoint: parse flags, read input, render ASCII art, write output.
func main() {
	// define `--font` flag (banner name without .txt) and default to "standard"
	font := flag.String("font", "standard", "banner name: standard, shadow or thinkertoy (filename without .txt)")

	// define `--out` flag for optional output file; empty means write to stdout
	out := flag.String("out", "", "output file (optional, defaults to stdout)")

	// parse the command-line flags provided by the user
	flag.Parse()

	// prepare a variable to hold the input text
	var input string

	// if positional args are present, join them into the input string
	if flag.NArg() > 0 {
		// join all remaining non-flag arguments with spaces
		input = strings.Join(flag.Args(), " ")
		if input == `\n` {
			fmt.Println()
			return
		}
		input = strings.ReplaceAll(input, "\\n", "\n")
	} else {
		// otherwise read the whole stdin into memory
		b, err := io.ReadAll(os.Stdin)
		if err != nil {
			// print a helpful error to stderr and exit with non-zero status
			fmt.Fprintln(os.Stderr, "failed reading stdin:", err)
			os.Exit(1)
		}
		// convert bytes read from stdin into a string
		input = string(b)
	}

	// remove only a trailing newline from the input (but leave other whitespace intact)
	input = strings.TrimRight(input, "\n")

	// if after trimming there's no input, inform the user and exit
	if input == "" {
		fmt.Fprintln(os.Stderr, "no input provided; pass text as arguments or via stdin")
		os.Exit(2)
	}

	// validate the input according to the pipeline rules (length, allowed control chars, etc.)
	if err := pipeline.ValidateInput(input); err != nil {
		fmt.Fprintln(os.Stderr, "invalid input:", err)
		os.Exit(3)
	}

	// tokenize the validated input into rune-based tokens for rendering
	tokens := pipeline.Tokenize(input)

	// load the requested banner (font) by name from the `banners/` directory
	banner, err := pipeline.LoadBanner(*font)
	if err != nil {
		// if loading fails, report and exit
		fmt.Fprintln(os.Stderr, "failed loading banner:", err)
		os.Exit(4)
	}

	// render the tokens into ASCII-art lines using the loaded banner
	lines := pipeline.RenderLines(tokens, banner)

	// default writer is stdout
	var w io.Writer = os.Stdout

	// if an output filename was provided, create the file and use it as writer
	if *out != "" {
		f, err := os.Create(*out)
		if err != nil {
			// report the create error and exit
			fmt.Fprintln(os.Stderr, "failed creating output file:", err)
			os.Exit(5)
		}
		// ensure the file is closed when main returns
		defer f.Close()
		// set writer to the created file
		w = f
	}

	// write the rendered lines to the chosen writer (stdout or file)
	if err := pipeline.WriteOutput(lines, w); err != nil {
		fmt.Fprintln(os.Stderr, "failed writing output:", err)
		os.Exit(6)
	}
}
