// Package postgresstore -
package postgresstore

import (
	"fmt"

	"database/sql"
	"errors"

	_ "github.com/lib/pq" //nolint:golint
)

type pg struct {
	dbConn *sql.DB
}

var ErrParseDBURL = errors.New("")

func NewPGStore(connStr string) (*pg, error) {
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to parse DATABASE_URL error %w", err)
	}

	if err != nil {
		return nil, fmt.Errorf("unable to create db connection with error %w", err)
	}

	p := pg{
		dbConn: conn,
	}

	return &p, nil
}

// CheckExists - check if a release exists with the nominated title.
func (pg *pg) CheckExists(feed, title string) (bool, error) {
	var tmpTitle sql.NullString

	err := pg.dbConn.QueryRow("select title from releases where feed = $1 and title = $2", feed, title).Scan(&tmpTitle)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false, fmt.Errorf("unable to query releases with error %w", err)
	}

	return tmpTitle.Valid, nil
}

// CreateItem - create a new release item in the database.
func (pg *pg) CreateItem(feed, title, content, link string) error {
	var id sql.NullInt64

	sqlStatement := `
	INSERT INTO releases (feed, title, content, link)
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	return pg.dbConn.QueryRow(sqlStatement, feed, title, content, link).Scan(&id)
}
