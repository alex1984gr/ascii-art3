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

	out := strings.Join(lines, "\n") + "\n"
	_, err := io.WriteString(w, out)
	return err
}
