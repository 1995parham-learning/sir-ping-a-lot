package migrate

import (
	"saver/store"
	"saver/store/status"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func Register(root *cobra.Command, d *gorm.DB) {
	c := cobra.Command{
		Use:   "migrate",
		Short: "Manages database, creates and fills tables if don't exist",
		Run: func(cmd *cobra.Command, args []string) {
			user := store.NewUser(d)
			user.Create()

			url := store.NewURL(d)
			url.Create()

			sqlStatus := status.NewSQLStatus(d)
			sqlStatus.Create()
		},
	}

	root.AddCommand(
		&c,
	)
}
