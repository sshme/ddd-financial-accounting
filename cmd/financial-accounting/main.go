package main

import (
	"ddd-financial-accounting/internal"
	"ddd-financial-accounting/internal/interface/cli"
)

func main() {
	var err error
	app, err := internal.InitializeApp()
	if err != nil {
		app.Logger.Error("Error initializing app", "error", err)
		return
	}

	p := cli.NewProgram(app)
	if _, err := p.Run(); err != nil {
		app.Logger.Error("Error running program", "error", err)
	}
}
