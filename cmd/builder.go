package cmd

type CommandBuilder struct {
}

func (builder *CommandBuilder) Build() *Cli {
	cli := NewCli()

	// set global flags for rootCmd in cli.
	cli.SetFlags()

	base := &baseCommand{cmd: cli.rootCmd}
	base.Cmd().SilenceErrors = true

	cli.AddCommand(base, &PlayCommand{})
	return cli
}
