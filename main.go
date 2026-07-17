package main

import (
	"fmt"
	"os"

	"github.com/tentharmy/rogue/inner/app"
	"github.com/tentharmy/rogue/inner/tui"
)

func main() {
	gameServ := app.NewGame()

	if err := tui.Run(gameServ); err != nil {
		fmt.Fprintf(os.Stderr, "Oof: %v\n", err)
		os.Exit(1)
	}
}
