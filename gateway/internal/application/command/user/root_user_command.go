package user

import (
	"github.com/spf13/cobra"
	"github.com/superdev/ecommerce/gateway/internal/command"
)

func NewRootUserCommand() command.CobraCommand {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "Command to create an manage users",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	return command.NewGenericCmd(cmd, "root:user")
}
