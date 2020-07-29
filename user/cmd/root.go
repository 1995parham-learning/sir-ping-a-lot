package cmd

import (
	"fmt"
	"os"
	"user/cmd/migrate"
	"user/cmd/server"
	"user/config"

	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// rootCmd represents the base command when called without any subcommands
	rootCmd := &cobra.Command{
		Use:   "user",
		Short: "A brief description of your application",
	}

	cfg := config.Read()

	migrate.Register(rootCmd, cfg.Database)
	server.Register(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
