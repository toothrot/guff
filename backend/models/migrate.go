package models

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/godoc_vfs"
	_ "github.com/golang-migrate/migrate/v4/source/godoc_vfs"
	_ "github.com/lib/pq"
	"golang.org/x/tools/godoc/vfs/mapfs"
)

var migrations = mapfs.New(map[string]string{
	"1_guff.up.sql": `CREATE TABLE divisions (
			id SERIAL,
            extid TEXT UNIQUE,
			name TEXT
	);`,
	"1_guff.down.sql": "DROP TABLE divisions",
	"2_users.up.sql": `CREATE TABLE users (
    	id SERIAL,
    	email TEXT UNIQUE,
    	is_admin BOOLEAN
	);`,
	"2_users.down.sql": `DROP TABLE users;`,
})

func Migrate(db *sql.DB) error {
	ms, err := godoc_vfs.WithInstance(migrations, "/")
	if err != nil {
		return err
	}
	pi, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithInstance("godoc_vfs", ms, "postgres", pi)
	if err != nil {
		return err
	}
	if err := m.Up(); err != migrate.ErrNoChange {
		return err
	}
	return nil
}
