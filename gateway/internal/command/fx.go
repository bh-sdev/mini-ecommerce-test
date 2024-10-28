package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var Module = fx.Module("command",
	// fx.WithLogger(func(log zap.Logger) fxevent.Logger {
	// 	return &fxevent.ZapLogger{Logger: log}
	// }),
	fx.Invoke(func(
		cmd *cobra.Command,
		shutdowner fx.Shutdowner,
	) error {
		if err := cmd.Execute(); err != nil {
			fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing the CLI '%s'", err)
			shutdowner.Shutdown(fx.ExitCode(1))
		}
		shutdowner.Shutdown()
		return nil

	}),
	fx.Provide(
		// AsCommand(NewHelloCmd),
		// AsCommand(NewHiCmd),
		fx.Annotate(
			NewRootCmd,
			fx.ParamTags(`group:"app.commands"`),
		),
	),
)
