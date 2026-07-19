package server

import (
	"github.com/httpmon/user/config"
	"github.com/httpmon/user/db"
	"github.com/httpmon/user/service"
	"github.com/httpmon/user/store"
	"github.com/spf13/cobra"
)

func Register(root *cobra.Command) {
	// nolint: exhaustivestruct
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
