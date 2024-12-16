package commands

import (
	"fmt"
	"strings"

	"github.com/c2micro/c2mcli/internal/service"
	"github.com/c2micro/c2mcli/internal/utils"
	"github.com/c2micro/c2mshr/defaults"
	"github.com/fatih/color"
	"github.com/reeflective/console"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func operatorListCommand(*console.Console) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list operators",
		Run: func(cmd *cobra.Command, _ []string) {
			operators, err := service.ListOperators()
			if err != nil {
				switch status.Code(err) {
				default:
					color.Red("list operators: %s", err.Error())
				}
				return
			}
			if len(operators) == 0 {
				color.Yellow("no operators exists")
				return
			}
			for _, o := range operators {
				fmt.Println(strings.Repeat("-", 42))
				fmt.Println(utils.PrintOperator(o))
			}
			fmt.Println(strings.Repeat("-", 42))
		},
	}
}

func operatorAddCommand(*console.Console) *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "add new operator",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			username := args[0]
			if len(username) < defaults.OperatorUsernameMinLength || len(username) > defaults.OperatorUsernameMaxLength {
				color.Red("invalid username")
				return
			}
			operator, err := service.AddOperator(username)
			if err != nil {
				switch status.Code(err) {
				case codes.AlreadyExists:
					color.Red("operator already exists")
				default:
					color.Red("add operator: %s", err.Error())
				}
				return
			}
			fmt.Println(utils.PrintOperator(operator))
		},
	}
}

func operatorRevokeCommand(*console.Console) *cobra.Command {
	return &cobra.Command{
		Use:   "revoke",
		Short: "revoke operator's token",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			username := args[0]
			if len(username) < defaults.OperatorUsernameMinLength || len(username) > defaults.OperatorUsernameMaxLength {
				color.Red("invalid username")
				return
			}
			if err := service.RevokeOperator(username); err != nil {
				switch status.Code(err) {
				default:
					color.Red("revoke operator: %s", err.Error())
				}
			}
			color.Green("%s token revoked", username)
		},
	}
}

func operatorRegenCommand(*console.Console) *cobra.Command {
	return &cobra.Command{
		Use:   "regen",
		Short: "regen operator's token",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			username := args[0]
			if len(username) < defaults.OperatorUsernameMinLength || len(username) > defaults.OperatorUsernameMaxLength {
				color.Red("invalid username")
				return
			}
			operator, err := service.RegenOperator(username)
			if err != nil {
				switch status.Code(err) {
				default:
					color.Red("regen operator: %s", err.Error())
				}
				return
			}
			fmt.Println(utils.PrintOperator(operator))
		},
	}
}

func operatorCommand(c *console.Console) *cobra.Command {
	operatorCmd := &cobra.Command{
		Use:     "operator",
		Short:   "manage operators",
		GroupID: operatorGroupId,
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Usage()
		},
	}

	operatorCmd.AddCommand(
		operatorListCommand(c),
		operatorAddCommand(c),
		operatorRevokeCommand(c),
		operatorRegenCommand(c),
	)

	return operatorCmd
}
