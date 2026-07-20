package migrate

import (
	"log"

	"github.com/httpmon/user/config"
	"github.com/httpmon/user/db"
	"github.com/httpmon/user/model"
	"github.com/spf13/cobra"
)

// Register wires the "migrate" sub-command, which creates the database tables
// from the GORM models if they don't already exist.
func Register(root *cobra.Command, cfg config.Database) {
	// nolint: exhaustivestruct
	c := cobra.Command{
		Use:   "migrate",
		Short: "Manages database, creates tables if they don't exist",
		Run: func(cmd *cobra.Command, args []string) {
			conn := db.New(cfg)

			if err := conn.AutoMigrate(&model.User{}, &model.URL{}); err != nil {
				log.Fatal(err)
			}
		},
	}

	root.AddCommand(
		&c,
	)
}
