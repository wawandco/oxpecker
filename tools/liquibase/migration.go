package liquibase

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/wawandco/oxpecker/internal/log"
)

// Migration xml with liquibase format. A migration may be composed
// of multiple changesets.
type Migration struct {
	ChangeSets []ChangeSet `xml:"changeSet"`
}

// ChangeSet with SQL and Rollback instructions.
type ChangeSet struct {
	ID          string `xml:"id,attr"`
	Author      string `xml:"author,attr"`
	SQL         string `xml:"sql"`
	RollbackSQL string `xml:"rollback"`
}

// Execute a changeset takes the SQL part of the changeset and runs it.
func (cs ChangeSet) Execute(conn *pgx.Conn, file string) error {
	var err error
	ctx := context.Background()

	row := conn.QueryRow(context.Background(), `SELECT orderexecuted FROM databasechangelog ORDER BY dateexecuted desc`)
	var order int
	if err = row.Scan(&order); err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	var count int
	row = conn.QueryRow(ctx, `SELECT count(*) FROM databasechangelog WHERE id = $1`, cs.ID)
	if err = row.Scan(&count); err == nil && count > 0 {
		return nil
	}

	if err != nil {
		return err
	}

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, cs.SQL)
	if err != nil {
		return err
	}

	insertStmt := `
		INSERT 
		INTO databasechangelog (id, author, filename, dateexecuted, orderexecuted,exectype) 
		VALUES ($1, $2, $3, $4, $5, $6);
	`

	_, err = tx.Exec(ctx, insertStmt, cs.ID, cs.Author, file, time.Now(), order+1, "EXECUTED")
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	log.Infof("Executed `%v`.", cs.ID)
	order++

	return nil
}

// Rollback the changeset runs the Rollback section of the
// changeset.
func (cs ChangeSet) Rollback(conn *pgx.Conn) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}

	log.Infof("Rolling back %v.", cs.ID)
	_, err = tx.Exec(context.Background(), cs.RollbackSQL)
	if err != nil {
		return err
	}

	_, err = tx.Exec(context.Background(), `DELETE FROM databasechangelog WHERE id = $1`, cs.ID)
	if err != nil {
		return err
	}

	return tx.Commit(context.Background())
}
