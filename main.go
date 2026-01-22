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

// main is the program entrypoint: parse flags, read input, render ASCII art, and write output.
func main() {
	// Define the `--font` flag to specify which banner to use (standard, shadow, or thinkertoy).
	// Defaults to "standard" if not provided.
	font := flag.String("font", "standard", "banner name: standard, shadow or thinkertoy (filename without .txt)")

	// Define the `--out` flag for optional output file path.
	// Empty string means output will be written to stdout.
	out := flag.String("out", "", "output file (optional, defaults to stdout)")

	// Parse the command-line flags provided by the user.
	flag.Parse()

	// Prepare a variable to hold the input text to be rendered as ASCII art.
	var input string

	// Check if positional arguments are present after flags.
	if flag.NArg() > 0 {
		// Join all remaining non-flag arguments into a single string separated by spaces.
		input = strings.Join(flag.Args(), " ")
		// Handle the special case where the only input is the literal string "\n".
		if input == `\n` {
			fmt.Println()
			return
		}
		// Replace the literal string "\n" with actual newline characters.
		input = strings.ReplaceAll(input, "\\n", "\n")
	} else {
		// If no positional arguments, read the entire input from stdin into memory.
		b, err := io.ReadAll(os.Stdin)
		if err != nil {
			// Print error to stderr and exit with status 1 if reading fails.
			fmt.Fprintln(os.Stderr, "failed reading stdin:", err)
			os.Exit(1)
		}
		// Convert the bytes read from stdin into a string.
		input = string(b)
	}

	// Remove only trailing newlines from the input (but preserve other whitespace).
	input = strings.TrimRight(input, "\n")

	// If after trimming there's no input left, inform the user and exit with status 2.
	if input == "" {
		fmt.Fprintln(os.Stderr, "no input provided; pass text as arguments or via stdin")
		os.Exit(2)
	}

	// Validate the input according to the pipeline rules (maximum length, allowed control characters, etc.).
	if err := pipeline.ValidateInput(input); err != nil {
		fmt.Fprintln(os.Stderr, "invalid input:", err)
		os.Exit(3)
	}

	// Tokenize the validated input into individual rune-based tokens for rendering.
	tokens := pipeline.Tokenize(input)

	// Load the requested banner (font) by name from the `banners/` directory.
	banner, err := pipeline.LoadBanner(*font)
	if err != nil {
		// If banner loading fails, report the error to stderr and exit with status 4.
		fmt.Fprintln(os.Stderr, "failed loading banner:", err)
		os.Exit(4)
	}

	// Render the tokens into ASCII-art lines (8 rows per character) using the loaded banner.
	lines := pipeline.RenderLines(tokens, banner)

	// Default writer is stdout (standard output).
	var w io.Writer = os.Stdout

	// If an output filename was provided via the --out flag, create and use the file as writer.
	if *out != "" {
		f, err := os.Create(*out)
		if err != nil {
			// If file creation fails, report the error to stderr and exit with status 5.
			fmt.Fprintln(os.Stderr, "failed creating output file:", err)
			os.Exit(5)
		}
		// Ensure the file is closed when main returns (even if an error occurs).
		defer f.Close()
		// Set the writer to the created file instead of stdout.
		w = f
	}

	// Write the ASCII art lines to the chosen writer (stdout or file).
	// WriteOutput will join the lines with newlines and add proper formatting.
	if err := pipeline.WriteOutput(lines, w); err != nil {
		fmt.Fprintln(os.Stderr, "failed writing output:", err)
		os.Exit(6)
	}
}
