package models

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	testDB = "guff_test"
)

func truncateTables(ctx context.Context, db *sql.DB) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, "TRUNCATE TABLE divisions; TRUNCATE TABLE users;"); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func initTestDb(ctx context.Context, t *testing.T) (*sql.DB, func()) {
	t.Helper()
	db, err := sql.Open("postgres", fmt.Sprintf("dbname=%s", testDB))
	if err != nil {
		t.Fatalf("sql.Open(%q, %q) = _, %v, wanted no error", "postgres", fmt.Sprintf("dbname=%q", testDB), err)
	}
	cleanup := func() {
		if err := truncateTables(ctx, db); err != nil {
			t.Errorf("truncateTables(%v, %v) = %v, wanted no error", ctx, db, err)
		}
		if err := db.Close(); err != nil {
			t.Errorf("db.Close() = %v, wanted no error", err)
		}
	}
	if err := Migrate(db); err != nil {
		cleanup()
		t.Fatalf("Migrate(%v) = %v, wanted no error", db, err)
	}
	if err := truncateTables(ctx, db); err != nil {
		cleanup()
		t.Fatalf("truncateTables(%v, %v) = %v, wanted no error", ctx, db, err)
	}
	return db, cleanup
}

func TestDBPersist_UpsertDivisions(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db, cleanup := initTestDb(ctx, t)
	defer cleanup()

	d := &DBPersist{DB: db}

	cases := []struct {
		desc    string
		want    []Division
		input   []Division
		current []Division
	}{
		{
			desc:  "inserting into empty table",
			want:  []Division{{ID: "123", Name: "Barf"}},
			input: []Division{{ID: "123", Name: "Barf"}},
		},
		{
			desc:    "inserting all duplicates",
			want:    []Division{{ID: "123", Name: "Barf"}},
			current: []Division{{ID: "123", Name: "Barf"}},
			input:   []Division{{ID: "123", Name: "Barf"}},
		},
		{
			desc:    "inserting an empty set",
			want:    []Division{{ID: "123", Name: "Barf"}},
			current: []Division{{ID: "123", Name: "Barf"}},
			input:   []Division{},
		},
		{
			desc:    "updating name",
			want:    []Division{{ID: "123", Name: "Beautiful People"}},
			current: []Division{{ID: "123", Name: "Barf"}},
			input:   []Division{{ID: "123", Name: "Beautiful People"}},
		},
	}
	for _, c := range cases {
		if err := truncateTables(ctx, db); err != nil {
			t.Fatalf("%q: truncateTables(%v, %v) = %v, wanted no error", c.desc, ctx, db, err)
		}

		// preload table with divisions for duplication tests
		if err := d.UpsertDivisions(ctx, c.current); err != nil {
			t.Errorf("%q: PersistDivisions(%v) = %v, wanted no error", c.desc, c.current, err)
			continue
		}
		if err := d.UpsertDivisions(ctx, c.input); err != nil {
			t.Errorf("%q: PersistDivisions(%v) = %v, wanted no error", c.desc, c.input, err)
			continue
		}

		rows, err := db.QueryContext(ctx, "SELECT extid, name FROM divisions")
		if err != nil {
			t.Errorf("%q: db.QueryContext(%v, %q) = _, %v, wanted no error", c.desc, ctx, "SELECT extid, name FROM divisions", err)
			continue
		}
		defer rows.Close()

		var got []Division
		for rows.Next() {
			var d Division
			if err := rows.Scan(&d.ID, &d.Name); err != nil {
				t.Fatalf("%q: rows.Scan(%v) = %v, wanted no error", c.desc, d, err)
			}
			got = append(got, d)
		}

		if diff := cmp.Diff(c.want, got); diff != "" {
			t.Errorf("%q: PersistDivisions() mismatch (-want +got):\n%s", c.desc, diff)
		}
	}
}

func TestDBPersist_GetDivisions(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db, cleanup := initTestDb(ctx, t)
	defer cleanup()

	query := "INSERT INTO divisions(extid, name) VALUES ('testid1', 'testname1')"
	if _, err := db.ExecContext(ctx, query); err != nil {
		t.Fatalf("db.ExecContext(%v, %q) = %v, wanted no error", ctx, query, err)
	}
	d := &DBPersist{DB: db}
	want := []Division{{ID: "testid1", Name: "testname1"}}

	got, err := d.GetDivisions(ctx)
	if err != nil {
		t.Errorf("d.GetDivisions(%v) = %v, wanted no error", ctx, err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("PersistDivisions() mismatch (-want +got):\n%s", diff)
	}
}

func TestDBPersist_FindOrCreateUser(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db, cleanup := initTestDb(ctx, t)
	defer cleanup()
	d := &DBPersist{
		DB: db,
	}

	tests := []struct {
		name    string
		email   string
		want    User
		wantErr bool
	}{
		{
			name:    "empty email",
			email:   " ",
			wantErr: true,
		},
		{
			name:  "first user, does not exist",
			email: "testuser1@example.com",
			want:  User{Email: "testuser1@example.com", IsAdmin: true},
		},
		{
			name:  "first user, created by previous test",
			email: "testuser1@example.com",
			want:  User{Email: "testuser1@example.com", IsAdmin: true},
		},
		{
			name:  "second user, does not get admin",
			email: "testuser2@example.com",
			want:  User{Email: "testuser2@example.com"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := d.FindOrCreateUser(ctx, tt.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBPersist.FindOrCreateUser(_, %v) error = %v, wantErr %v", tt.email, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FindOrCreateUser(_, %v) mismatch (-want +got):\n%s", tt.email, diff)
			}
		})
	}
}
