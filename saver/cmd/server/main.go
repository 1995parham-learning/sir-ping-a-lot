package server

import (
	"saver/config"

	"github.com/nats-io/go-nats"
	"github.com/spf13/cobra"
)

func Register(root *cobra.Command, cfg config.Nats) {
	c := cobra.Command{
		Use:   "server",
		Short: "Manages database, creates and fills tables if don't exist",
		Run: func(cmd *cobra.Command, args []string) {
			n := balancer.New(cfg.Nats)
		},
	}

	root.AddCommand(
		&c,
	)
}
