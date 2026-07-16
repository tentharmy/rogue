package main

import (
	"fmt"
	"os"

	"github.com/tentharmy/rogue/inner/tui"
)

func main() {
	if err := tui.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Oof: %v\n", err)
		os.Exit(1)
	}
}
