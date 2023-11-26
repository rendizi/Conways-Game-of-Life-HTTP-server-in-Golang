package main

import (
	"context"
	"github.com/aivanov/game/internal/application"
	"os"
)

func main() {
	ctx := context.Background()
	// Exit leads to the termination of the program with the given code.
	os.Exit(mainWithExitCode(ctx))
}

func mainWithExitCode(ctx context.Context) int {
	cfg := application.Config{
		Width:  50,
		Height: 50,
	}
	app := application.New(cfg)

	return app.Run(ctx)
}
