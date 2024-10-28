package user

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/superdev/ecommerce/gateway/internal/command"
	"github.com/superdev/ecommerce/gateway/internal/domain/entity"
	"github.com/superdev/ecommerce/gateway/internal/domain/service"
)

func NCreateUserCommand(
	userService service.UserService,
) command.CobraCommand {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create new user",
		Run: func(cmd *cobra.Command, args []string) {
			// cmd.Help()
			user, err := CreateUser(cmd, args, userService)
			if err != nil {
				cmd.PrintErrf("Could not create user: %s", err.Error())
			} else {
				cmd.PrintErrf("Created user: %s with id: %s", user.GetUserName(), user.GetId())
			}
		},
	}
	return command.NewGenericCmd(cmd, "root:user:create")
}

func CreateUser(cmd *cobra.Command, args []string, userService service.UserService) (entity.User, error) {

	var user = entity.NewEmptyUser()

	confirmForm := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("UserName").
				Placeholder("Enter Username").
				Description("Username for Login").
				Value(&user.Username),

			huh.NewInput().
				Title("User Password").
				Placeholder("Enter Password").
				Description("Password for Login").
				Value(&user.Password).
				EchoMode(huh.EchoModePassword),
			huh.NewSelect[string]().
				Title("User Role").
				Options(
					huh.NewOption("User", "user"),
					huh.NewOption("Admin", "admin"),
				).
				Value(&user.Role),
		),
	)
	err := confirmForm.Run()
	if err != nil {
		return nil, err
	}

	fmt.Printf("Saving User: %v\n", user)
	fmt.Printf("Saving Username: %v\n", user.GetUserName())

	_user, err := userService.Save(user)
	if err != nil {
		return nil, err
	}

	return _user, err
}
