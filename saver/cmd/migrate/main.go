package migrate

import (
	"database/sql"
	"log"
	"saver/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func Register(root *cobra.Command, d *gorm.DB, cfg config.Database) {
	c := cobra.Command{
		Use:   "migrate",
		Short: "Manages database, creates and fills tables if don't exist",
		Run: func(cmd *cobra.Command, args []string) {
			db, err := sql.Open("postgres", cfg.Cstring())
			if err != nil {
				log.Fatal(err)
			}

			driver, err := postgres.WithInstance(db, &postgres.Config{})
			if err != nil {
				log.Fatal(err)
			}

			p, err := migrate.NewWithDatabaseInstance("file://./migration", "monitor", driver)
			if err != nil {
				log.Fatal(err)
			}

			if err := p.Up(); err != nil && err != migrate.ErrNoChange {
				log.Fatal(err)
			}
		},
	}

	root.AddCommand(
		&c,
	)
}
