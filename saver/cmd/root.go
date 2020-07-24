package cmd

import (
	"HTTP_monitoring/memory"
	"fmt"
	"os"
	"saver/balancer"
	"saver/cmd/migrate"
	"saver/cmd/server"
	"saver/config"
	"saver/db"
	"saver/store/status"

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
	r := memory.New(cfg.Redis)
	n := balancer.New(cfg.Nats)

	migrate.Register(rootCmd, d)
	server.Register(rootCmd, n, cfg.Nats, status.NewRedisStatus(r), cfg.Redis,  status.NewSQLStatus(d))

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
