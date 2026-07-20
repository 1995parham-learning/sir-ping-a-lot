package migrate

import (
	"log"
	"saver/config"
	"saver/db"
	"saver/model"

	"github.com/spf13/cobra"
)

// expiryTrigger keeps the statuses table small by deleting rows older than two
// days after every insert. GORM's AutoMigrate can't express triggers, so we run
// this alongside it with a plain Exec.
const expiryTrigger = `
CREATE OR REPLACE FUNCTION delete_expired_row()
RETURNS TRIGGER AS
	$BODY$
		BEGIN
		DELETE FROM statuses WHERE clock < NOW() - INTERVAL '2 days';
		RETURN NULL;
		END;
	$BODY$
	LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS delete_expired_rows ON statuses;

CREATE TRIGGER delete_expired_rows
	AFTER INSERT ON statuses
	FOR EACH ROW
	EXECUTE PROCEDURE delete_expired_row();
`

// Register wires the "migrate" sub-command, which creates the database tables
// from the GORM models and installs the status-expiry trigger.
func Register(root *cobra.Command, cfg config.Database) {
	// nolint: exhaustivestruct
	c := cobra.Command{
		Use:   "migrate",
		Short: "Manages database, creates tables if they don't exist",
		Run: func(cmd *cobra.Command, args []string) {
			conn := db.New(cfg)

			if err := conn.AutoMigrate(&model.User{}, &model.URL{}, &model.Status{}); err != nil {
				log.Fatal(err)
			}

			if err := conn.Exec(expiryTrigger).Error; err != nil {
				log.Fatal(err)
			}
		},
	}

	root.AddCommand(
		&c,
	)
}
