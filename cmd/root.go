package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/prs-io/openobserve-go/api"
	"github.com/prs-io/openobserve-go/client"

	"github.com/spf13/cobra"
)

var (
	orgID    string
	host     string
	username string
	password string
)

var API *api.API

// https://github.com/kubernetes/kubectl/blob/88bb12ba04025e5b03a141ad26fca1c735d54e11/pkg/util/util.go#L33
// rootCmd represents the base command when called without any subcommands

func Execute() {
	cobra.CheckErr(rootCmd.ExecuteContext(context.Background()))
}

var rootCmd = &cobra.Command{
	Use:   "openobserve",
	Short: "OpenObserve CLI",
	Long:  "A command-line interface for interacting with OpenObserve.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Host: %s\n", host)
		fmt.Printf("Org ID: %s\n", orgID)
		if orgID == "" {
			cobra.CheckErr("--org flag is required")
		}
		if host == "" {
			cobra.CheckErr("--host flag is required")
		}

		if username = os.Getenv("OPENOBSERVE_USERNAME"); username == "" {
			cobra.CheckErr("envionment variable 'OPENOBSERVE_USERNAME' is not set")
		}
		if password = os.Getenv("OPENOBSERVE_PASSWORD"); password == "" {
			cobra.CheckErr("envionment variable 'OPENOBSERVE_PASSWORD' is not set")
		}

		cfg := client.Config{
			BaseURL:  host,
			Username: username,
			Password: password,
			Timeout:  10 * time.Second,
		}
		fmt.Printf("Config: %+v\n", cfg)
		openObserveClient := client.NewClient(cfg)
		API = api.NewAPI(openObserveClient)
		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&orgID, "org", "", "Organization ID (required)")
	rootCmd.PersistentFlags().StringVar(&host, "host", "", "OpenObserve host URL (required)")
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(usersCmd)

}

func bootStrap() {

	// cfg := client.Config{

}
