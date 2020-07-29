package server

import (
	"user/config"
	"user/db"
	"user/service"
	"user/store"

	"github.com/spf13/cobra"
)

func Register(root *cobra.Command) {
	c := cobra.Command{
		Use:   "server",
		Short: "Runs endpoints",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := config.Read()
			d := db.New(cfg.Database)

			api := service.API{
				User:   store.NewUser(d),
				URL:    store.NewURL(d),
				Config: cfg.JWT,
			}

			api.Run()
		},
	}

	root.AddCommand(
		&c,
	)
}