package cmd

import (
	"HTTP_monitoring/balancer"
	"HTTP_monitoring/memory"
	"fmt"
	"os"
	"saver/cmd/migrate"
	"saver/config"
	"saver/db"

	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:   "saver",
		Short: "A brief description of your application",
		// Uncomment the following line if your bare application
		// has an action associated with it:
		//	Run: func(cmd *cobra.Command, args []string) { },
	}

	cfg := config.Read()
	d := db.New(cfg.Database)

	migrate.Register(rootCmd, d)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
