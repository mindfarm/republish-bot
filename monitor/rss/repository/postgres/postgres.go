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

// NewPGStore -
// nolint:golint
func NewPGStore(connStr string) (*pg, error) {
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse DATABASE_URL error %v", err)
	}

	if err != nil {
		return nil, fmt.Errorf("unable to create db connection with error %w", err)
	}
	p := pg{
		dbConn: conn,
	}
	return &p, nil
}

// CheckExists - check if a release exists with the nominated title
func (pg *pg) CheckExists(title string) (bool, error) {
	var tmpTitle sql.NullString
	err := pg.dbConn.QueryRow("select title from releases where title = $1", title).Scan(&tmpTitle)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false, fmt.Errorf("unable to query releases with error %w", err)
	}
	return tmpTitle.Valid, nil

}

// CreateItem - create a new release item in the database
func (pg *pg) CreateItem(title, content, link string) error {
	var id sql.NullInt64
	sqlStatement := `
	INSERT INTO releases (title, content, link)
	VALUES ($1, $2, $3)
	RETURNING id`
	return pg.dbConn.QueryRow(sqlStatement, title, content, link).Scan(&id)
}
