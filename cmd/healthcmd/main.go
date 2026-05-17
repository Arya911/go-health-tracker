package main

import (
	"github.com/Arya911/healthCli/internal/checker"
	"github.com/Arya911/healthCli/internal/tui"
)

func main() {
	stream := checker.Start()
	app := tui.New(stream)

	app.Run()
}
