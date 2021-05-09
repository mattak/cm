package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	KickCmd = &cobra.Command{
		Use:   "kick [PROJECT_ID]",
		Short: "Kick cloud message",
		Long:  "Kick cloud message",
		Example: `
  cm kick myproject < reservation.jsons

reservation.jsons format: 
  {"token":"token1","data":{}}
  {"token":"token2","data":{}}
`,
		Run: runCommandKick,
	}
	isDryRun = false
)

func init() {
	KickCmd.PersistentFlags().BoolVarP(&isDryRun, "dryrun", "d", false, "dry run to notifications")
}

func runCommandKick(cmd *cobra.Command, args []string) {
	_, exists := os.LookupEnv("GOOGLE_APPLICATION_CREDENTIALS")
	if !exists {
		log.Fatal("Please set environment variable: GOOGLE_APPLICATION_CREDENTIALS")
	}
	if len(args) != 1 {
		log.Fatal("argument missing: projectId")
	}
	projectId := args[0]

	lines := ReadStdinLines()
	msgs, err := ReadMessageFromStdinJson(lines)
	if err != nil {
		log.Fatal(err)
	}

	PushNotifications(projectId, msgs, isDryRun)
}
