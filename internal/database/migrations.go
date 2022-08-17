package database

import (
	"database/sql"
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"github.com/jmoiron/sqlx"
	"github.com/lncapital/torq/database/migrations"
	"github.com/lncapital/torq/internal/channels"
	"log"
	"net/http"
)

// newMigrationInstance fetches sql files and creates a new migration instance.
func newMigrationInstance(db *sql.DB) (*migrate.Migrate, error) {
	sourceInstance, err := httpfs.New(http.FS(migrations.MigrationFiles), ".")
	if err != nil {
		return nil, fmt.Errorf("invalid source instance, %w", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithInstance("httpfs", sourceInstance, "postgres", driver)
	if err != nil {
		return nil, fmt.Errorf("could not create migration instance: %v", err)
	}

	return m, nil
}

// MigrateUp migrates up to the latest migration version. It should be used when the version number changes.
func MigrateUp(db *sqlx.DB) error {
	m, err := newMigrationInstance(db.DB)
	if err != nil {
		return err
	}

	version, _, err := m.Version()
	if err != nil {
		if !errors.As(err, &migrate.ErrNilVersion) {
			return errors.Wrap(err, "Getting database migration version")
		}
	}

	runMigration38 := func() error {
		if err = convertShortChannelIds(db); err != nil {
			return errors.Wrap(err, "Converting short channel ids")
		}
		m.Force(38)
		if err != nil {
			return errors.Wrap(err, "Setting database migration version to 38")
		}
		return nil
	}
	// three migrations for converting between c-lightning and lnd short channel id formats
	if version < 37 {
		err = m.Migrate(37)
		if err != nil {
			return errors.Wrap(err, "Migrating to version 37")
		}
		if err = runMigration38(); err != nil {
			return err
		}
	}

	// this state should be impossible but could happen if the migration process is interrupted
	if version == 37 {
		if err = runMigration38(); err != nil {
			return err
		}
	}

	err = m.Up()
	if err != nil {
		return errors.Wrap(err, "Migrating database up")
	}

	return nil
}

func convertShortChannelIds(db *sqlx.DB) error {

	log.Println("Running short channel id conversions")
	log.Println("WARNING: This process can take 10+ minutes")
	log.Println("Please do not interrupt")

	// channel table
	{
		log.Println("Step: 1 of 5: Updating channel table. ")
		rows, err := db.Query("SELECT short_channel_id FROM channel;")
		if err != nil {
			return errors.Wrap(err, "Selecting short_channel_id from channel")
		}

		for rows.Next() {
			var shortChannelId string
			err = rows.Scan(&shortChannelId)
			if err != nil {
				return errors.Wrap(err, "Scanning short channel id from channel table")
			}
			lndShortChannelId, err := channels.ConvertShortChannelIDToLND(shortChannelId)
			if err != nil {
				return errors.Wrap(err, "Converting short channel id to LND format")
			}
			updateStatement := "UPDATE channel SET lnd_short_channel_id = $1 WHERE short_channel_id = $2"
			if _, err := db.Exec(updateStatement, lndShortChannelId, shortChannelId); err != nil {
				return errors.Wrap(err, "Updating lnd_short_channel_id on channel table")
			}
		}
		err = rows.Err()
		if err != nil {
			return errors.Wrap(err, "Iterating over each channel row")
		}
	}

	// channel_event table
	{
		log.Println("Step: 2 of 5: Updating channel_event table. ")
		rows, err := db.Query(`
			SELECT distinct on (lnd_short_channel_id) lnd_short_channel_id as lnd_short_channel_id
			FROM channel_event;
		`)

		if err != nil {
			return errors.Wrap(err, "error selecting short_channel_id from channel_event table")
		}

		for rows.Next() {
			var lndShortChannelId uint64
			err = rows.Scan(&lndShortChannelId)
			if err != nil {
				return errors.Wrap(err, "Scanning short_channel_id and from channel_event table")
			}
			shortChannelId := channels.ConvertLNDShortChannelID(lndShortChannelId)

			updateStatement := `UPDATE channel_event SET short_channel_id = $1 WHERE
									lnd_short_channel_id = $2`
			if _, err := db.Exec(updateStatement, shortChannelId, lndShortChannelId); err != nil {
				return errors.Wrapf(err,
					"updating short_channel_id on channel_event table failed at: %v", lndShortChannelId)
			}

		}
		err = rows.Err()
		if err != nil {
			return errors.Wrap(err, "Iterating over each channel_event row")
		}
	}

	// forward table
	{
		log.Println("Step: 3 of 5: Updating forward table.")
		rows, err := db.Query(`
			select distinct on (lnd_short_channel_id) lnd_short_channel_id from(
				SELECT distinct on (lnd_outgoing_short_channel_id) lnd_outgoing_short_channel_id as lnd_short_channel_id
				FROM forward
				UNION
				SELECT distinct on (lnd_incoming_short_channel_id) lnd_incoming_short_channel_id as lnd_short_channel_id
				FROM forward) a
			where lnd_short_channel_id != 0;
		`)

		if err != nil {
			return errors.Wrap(err, "error selecting short_channel_id from forward table")
		}

		for rows.Next() {
			var lndShortChannelId uint64
			err = rows.Scan(&lndShortChannelId)
			if err != nil {
				return errors.Wrap(err, "Scanning short_channel_id and from forward table")
			}
			shortChannelId := channels.ConvertLNDShortChannelID(lndShortChannelId)

			updateOutgoingStatement := `UPDATE forward SET outgoing_short_channel_id = $1 WHERE
									lnd_outgoing_short_channel_id = $2`
			if _, err := db.Exec(updateOutgoingStatement, shortChannelId, lndShortChannelId); err != nil {
				return errors.Wrapf(err,
					"updating outgoing_short_channel_id on forward table failed at: %v", lndShortChannelId)
			}

			updateIncomingStatement := `UPDATE forward SET incoming_short_channel_id = $1 WHERE
									lnd_incoming_short_channel_id = $2`
			if _, err := db.Exec(updateIncomingStatement, shortChannelId, lndShortChannelId); err != nil {
				return errors.Wrapf(err,
					"updating incoming_short_channel_id on forward table failed at: %v", lndShortChannelId)
			}

		}
		err = rows.Err()
		if err != nil {
			return errors.Wrap(err, "Iterating over each forward table row")
		}
	}

	// htlc_event table
	{
		log.Println("Step: 4 of 5: Updating htlc_event table.")
		rows, err := db.Query(`
			select distinct on (lnd_short_channel_id) lnd_short_channel_id from(
				SELECT distinct on (lnd_outgoing_short_channel_id) lnd_outgoing_short_channel_id as lnd_short_channel_id
				FROM htlc_event
				where lnd_outgoing_short_channel_id != 0
				UNION
				SELECT distinct on (lnd_incoming_short_channel_id) lnd_incoming_short_channel_id as lnd_short_channel_id
				FROM htlc_event
				where lnd_incoming_short_channel_id != 0
			order by lnd_short_channel_id) a;
		`)
		if err != nil {
			return errors.Wrap(err, "Selecting lnd_outgoing_short_channel_id and lnd_incoming_short_channel_id from htlc_event")
		}

		for rows.Next() {
			var lndShortChannelId uint64
			err = rows.Scan(&lndShortChannelId)
			if err != nil {
				return errors.Wrap(err, "Scanning lnd_outgoing_short_channel_id and lnd_incoming_short_channel_id from htlc_event table")
			}

			shortChannelId := channels.ConvertLNDShortChannelID(lndShortChannelId)

			updateOutgoingStatement := `UPDATE htlc_event SET outgoing_short_channel_id = $1 WHERE
									lnd_outgoing_short_channel_id = $2`
			if _, err := db.Exec(updateOutgoingStatement, shortChannelId, lndShortChannelId); err != nil {
				return errors.Wrapf(err,
					"updating outgoing_short_channel_id on htlc_event table failed at: %v", lndShortChannelId)
			}

			updateIncomingStatement := `UPDATE htlc_event SET incoming_short_channel_id = $1 WHERE
									lnd_incoming_short_channel_id = $2`
			if _, err := db.Exec(updateIncomingStatement, shortChannelId, lndShortChannelId); err != nil {
				return errors.Wrapf(err,
					"updating incoming_short_channel_id on htlc_event table failed at: %v", lndShortChannelId)
			}
		}
		err = rows.Err()
		if err != nil {
			return errors.Wrap(err, "Iterating over each htlc_event table row")
		}
	}

	// routing_policy table
	{
		log.Println("Step: 5 of 5: Updating routing_policy table.")
		rows, err := db.Query(`
			SELECT distinct on (lnd_short_channel_id) lnd_short_channel_id as lnd_short_channel_id
			FROM routing_policy;
		`)

		if err != nil {
			return errors.Wrap(err, "error selecting short_channel_id from routing_policy table")
		}

		for rows.Next() {
			var lndShortChannelId uint64
			err = rows.Scan(&lndShortChannelId)
			if err != nil {
				return errors.Wrap(err, "Scanning short_channel_id and from routing_policy table")
			}
			shortChannelId := channels.ConvertLNDShortChannelID(lndShortChannelId)

			updateStatement := `UPDATE routing_policy SET short_channel_id = $1 WHERE
									lnd_short_channel_id = $2`
			if _, err := db.Exec(updateStatement, shortChannelId, lndShortChannelId); err != nil {
				return errors.Wrapf(err,
					"updating short_channel_id on routing_policy table failed at: %v", lndShortChannelId)
			}

		}
		err = rows.Err()
		if err != nil {
			return errors.Wrap(err, "Iterating over each routing_policy row")
		}
	}

	return nil
}
