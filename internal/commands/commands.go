package commands

import (
	"github.com/reeflective/console"
	"github.com/spf13/cobra"
)

func Commands(app *console.Console) console.Commands {
	return func() *cobra.Command {
		rootCmd := &cobra.Command{
			DisableFlagsInUseLine: true,
		}
		// регистрация групп
		rootCmd.AddGroup(
			&cobra.Group{ID: globalGroupId, Title: globalGroupId},
			&cobra.Group{ID: operatorGroupId, Title: operatorGroupId},
			&cobra.Group{ID: listenerGroupId, Title: listenerGroupId},
		)
		// exit
		rootCmd.AddCommand(exitCommand(app))
		// operator
		rootCmd.AddCommand(operatorCommand(app))
		// listener
		rootCmd.AddCommand(listenerCommand(app))
		
		rootCmd.SetHelpCommandGroupID(globalGroupId)
		rootCmd.InitDefaultHelpCmd()
		rootCmd.CompletionOptions.DisableDefaultCmd = true
		return rootCmd
	}
}
