package migrate

import (
	"database/sql"
	"errors"
	"log"
	"saver/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"// Imported for its side effects
	"github.com/spf13/cobra"
)

func Register(root *cobra.Command, cfg config.Database) {
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

			if err := p.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
				log.Fatal(err)
			}
		},
	}

	root.AddCommand(
		&c,
	)
}
