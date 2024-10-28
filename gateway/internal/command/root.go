package command

import (
	"strings"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func NewRootCmd(
	commands []CobraCommand,
	lc fx.Lifecycle,
	shutdowner fx.Shutdowner,
	// runner service_runners.ServerCommandRunner,
) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "root",
		Short: "SuperDev eCommerce Command-line Interface",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			// runner(cmd, context.Background())
		},
	}

	// message := fmt.Sprintf("COMMANDS: %v", commands)
	// fmt.Println(message)

	// Loop through commands adding them to other commands
	done := make(map[string]bool)
	for _, c1 := range commands {
		for _, c2 := range commands {
			if c1.FullName() == c2.FullName() || done[c1.FullName()] {
				continue
			}
			fullParent := strings.Join(strings.Split(c1.FullName(), ":")[:len(strings.Split(c1.FullName(), ":"))-1], ":")
			if fullParent == c2.FullName() {
				// Add the child to the parent
				c2.Command().AddCommand(c1.Command())
				done[c1.FullName()] = true
				continue
			}
			if fullParent == "root" {
				cmd.AddCommand(c1.Command())
				done[c1.FullName()] = true
				continue
			}
		}
	}

	// lc.Append(fx.Hook{
	// 	OnStart: func(ctx context.Context) error {
	// 		// log.Debug().Msg("\n Root Command OnStarted LifeCycle Called\n")
	// 		if err := cmd.Execute(); err != nil {
	// 			fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing the CLI '%s'", err)
	// 			shutdowner.Shutdown(fx.ExitCode(1))
	// 		}
	// 		shutdowner.Shutdown()
	// 		return nil
	// 	},
	// })
	return cmd
}
