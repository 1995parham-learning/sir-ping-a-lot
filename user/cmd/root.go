package cmd

import (
	"log"
	"os"

	"github.com/httpmon/user/cmd/migrate"
	"github.com/httpmon/user/cmd/server"
	"github.com/httpmon/user/config"
	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// rootCmd represents the base command when called without any subcommands
	// nolint: exhaustivestruct
	rootCmd := &cobra.Command{
		Use:   "github.com/httpmon/user",
		Short: "A brief description of your application",
	}

	cfg := config.Read()

	migrate.Register(rootCmd, cfg.Database)
	server.Register(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
