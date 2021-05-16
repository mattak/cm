package cmd

import "github.com/spf13/cobra"

var (
	RootCmd = &cobra.Command{
		Use:     "cm",
		Short:   "Cloud message delivery",
		Long:    "Cloud message delivery",
		Version: VERSION,
	}
)

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	RootCmd.AddCommand(KickCmd)
	RootCmd.AddCommand(VersionCmd)
}
