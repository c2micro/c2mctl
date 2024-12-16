package console

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/c2micro/c2mcli/internal/commands"
	"github.com/c2micro/c2mcli/internal/service"
	"github.com/c2micro/c2mcli/internal/utils"
	"github.com/fatih/color"
	"github.com/reeflective/console"
)

func Run(ctx context.Context) error {
	app := console.New("c2mctl")
	main := app.ActiveMenu()
	main.Short = "management commands"
	main.Prompt().Primary = func() string { return fmt.Sprintf("[%s] > ", color.CyanString("c2mctl")) }
	main.SetCommands(commands.Commands(app))
	main.AddInterrupt(io.EOF, func(c *console.Console) {
		if utils.ExitConsole(c) {
			service.Close()
			os.Exit(0)
		}
	})
	return app.StartContext(ctx)
}
