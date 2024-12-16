package cmd

import (
	"fmt"

	"github.com/c2micro/c2mcli/internal/console"
	"github.com/c2micro/c2mcli/internal/constants"
	"github.com/c2micro/c2mcli/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type App struct {
	Host  string
	Token string
}

func (a *App) RegisterFlags(f *pflag.FlagSet) {
	f.StringVarP(&a.Host, "host", "H", utils.EnvOr(constants.CmdEnvHostKey, ""), "host for management server")
	f.StringVarP(&a.Token, "token", "t", utils.EnvOr(constants.CmdEnvTokenKey, ""), "management token")
}

func (a *App) Validate() error {
	if a.Host == "" {
		return fmt.Errorf("host required")
	}
	if a.Token == "" {
		return fmt.Errorf("token required")
	}
	return nil
}

func (a *App) Run(cmd *cobra.Command, _ []string) error {
	return console.Run(cmd.Context())
}
