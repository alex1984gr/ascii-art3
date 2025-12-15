package pipeline

import (
	"errors"
	"io"
	"strings"
)

func WriteOutput(lines []string, w io.Writer) error {
	if w == nil {
		return errors.New("writer is nil")
	}

	if len(lines) == 0 {
		return nil
	}

	lines[0] = strings.TrimLeft(lines[0], " ")

	for _, line := range lines {
		// Αν η γραμμή είναι literal "\n", τυπώνουμε πραγματικό newline
		if line == `\n` {
			_, err := io.WriteString(w, "\n")
			if err != nil {
				return err
			}
			continue
		}

		_, err := io.WriteString(w, line+"\n")
		if err != nil {
			return err
		}
	}

	return nil
}
