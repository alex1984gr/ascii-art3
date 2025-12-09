package main

import (
	"ascii-art/pipeline"
	"fmt"
)

func main() {
	input := "Hello There"
	tokens := pipeline.Tokenize(input)
	fmt.Println("Tokens:", tokens)
	banner, err := pipeline.LoadBanner("standard")
	if err != nil {
		fmt.Println("LoadBanner error:", err)
		return
	}
	for i, t := range tokens {
		g, ok := banner[t]
		if !ok {
			fmt.Printf("[%d] token=%q -> MISSING\n", i, t)
		} else {
			fmt.Printf("[%d] token=%q -> glyph rows[0..2]:\n", i, t)
			end := 3
			if len(g) < end {
				end = len(g)
			}
			for j := 0; j < end; j++ {
				fmt.Printf("    %q\n", g[j])
			}
		}
	}

	// Render and show the assembled output lines
	out := pipeline.RenderLines(tokens, banner)
	fmt.Println("\nRendered lines:")
	for i, l := range out {
		fmt.Printf("%2d: %q\n", i, l)
	}
}
