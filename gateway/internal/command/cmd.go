package command

import (
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

type CobraCommand interface {
	Command() *cobra.Command
	FullName() string
}

type GenericCommand struct {
	cmd      *cobra.Command
	fullName string
}

func (g GenericCommand) Command() *cobra.Command {
	return g.cmd
}

func (g GenericCommand) FullName() string {
	return g.fullName
}

// ---------------

func NewGenericCmd(
	cmd *cobra.Command,
	fullName string,
) GenericCommand {

	return GenericCommand{cmd: cmd, fullName: fullName}
}

// func AsCommand(f any) any {
// 	return fx.Annotate(
// 		f,
// 		fx.As(new(CobraCommand)),
// 		fx.ResultTags(`group:"commands"`),
// 	)
// }

func AsCobraCommand(f any, anns ...fx.Annotation) any {
	anns = append(anns,
		fx.As(new(CobraCommand)),
		fx.ResultTags(`group:"app.commands"`))

	return fx.Annotate(
		f,
		anns...,
	)
}
