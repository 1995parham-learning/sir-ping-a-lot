package server

import (
	"saver/config"
	"saver/store/status"
	"saver/subscriber"

	"github.com/nats-io/go-nats"
	"github.com/spf13/cobra"
)

func Register(root *cobra.Command, n *nats.EncodedConn, natsCfg config.Nats,
	r status.RedisStatus, redisCfg config.Redis, s status.SQLStatus) {
	c := cobra.Command{
		Use:   "server",
		Short: "Subscribes to the right topic",
		Run: func(cmd *cobra.Command, args []string) {
			server := subscriber.New(n, natsCfg, &r, redisCfg, s)
			server.Subscribe()
		},
	}

	root.AddCommand(
		&c,
	)
}
