package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"sync"
)

type Persist interface {
	GetDivisions(ctx context.Context) ([]Division, error)
	TruncateDivisions(ctx context.Context) error
	UpsertDivisions(ctx context.Context, ds []Division) error

	FindOrCreateUser(ctx context.Context, email string) (User, error)
	TruncateUsers(ctx context.Context) error
}

type DBPersist struct {
	DB *sql.DB
}

func (d *DBPersist) TruncateUsers(ctx context.Context) error {
	return errors.New("unimplemented")
}

func (d *DBPersist) FindOrCreateUser(ctx context.Context, email string) (User, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	email = strings.TrimSpace(email)
	if email == "" {
		return User{}, fmt.Errorf("invalid email %q", email)
	}

	u, err := d.getUser(ctx, email)
	if err != nil {
		if err != sql.ErrNoRows {
			return u, err
		}
		if u, err = d.createUser(ctx, email); err != nil {
			return u, err
		}
	}
	return u, nil
}

func (d *DBPersist) getUser(ctx context.Context, email string) (User, error) {
	var u User
	row := d.DB.QueryRowContext(ctx, "SELECT email, is_admin FROM users WHERE email = $1", email)
	err := row.Scan(&u.Email, &u.IsAdmin)
	return u, err
}

func (d *DBPersist) createUser(ctx context.Context, email string) (User, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var userCount int64
	u := User{Email: email}

	tx, err := d.DB.BeginTx(ctx, nil)
	if err != nil {
		return User{}, err
	}
	row := tx.QueryRowContext(ctx, "SELECT COUNT(*) FROM users")
	if err := row.Scan(&userCount); err != nil {
		return u, err
	}
	if userCount == 0 {
		u.IsAdmin = true
	}
	if _, err := tx.ExecContext(ctx, "INSERT INTO users(email, is_admin) VALUES ($1, $2) ON CONFLICT DO NOTHING", u.Email, u.IsAdmin); err != nil {
		return u, err
	}
	if err := tx.Commit(); err != nil {
		return u, err
	}
	return u, nil
}

func (d *DBPersist) TruncateDivisions(ctx context.Context) error {
	return errors.New("unimplemented")
}

func (d *DBPersist) GetDivisions(ctx context.Context) ([]Division, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	rows, err := d.DB.QueryContext(ctx, "SELECT extid, name FROM divisions;")
	if err != nil {
		return nil, err
	}
	var ds []Division
	for rows.Next() {
		var d Division
		if err := rows.Scan(&d.ID, &d.Name); err != nil {
			return ds, err
		}
		ds = append(ds, d)
	}
	return ds, nil
}

func (d *DBPersist) UpsertDivisions(ctx context.Context, ds []Division) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if len(ds) == 0 {
		return nil
	}

	tx, err := d.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	s, err := tx.PrepareContext(ctx, "INSERT INTO divisions(extid, name) VALUES ($1, $2) ON CONFLICT(extid) DO NOTHING")
	if err != nil {
		return err
	}
	for _, d := range ds {
		if _, err := s.ExecContext(ctx, d.ID, d.Name); err != nil {
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

var DefaultMemoryPersist = &MemoryPersist{
	divisions: []Division{},
	mx:        &sync.Mutex{},
	users:     map[string]User{},
}

type MemoryPersist struct {
	divisions []Division
	users     map[string]User
	mx        *sync.Mutex
}

func (m *MemoryPersist) TruncateUsers(ctx context.Context) error {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.users = map[string]User{}
	return nil
}

func (m *MemoryPersist) FindOrCreateUser(ctx context.Context, email string) (User, error) {
	m.mx.Lock()
	defer m.mx.Unlock()
	u, ok := m.users[email]
	if !ok {
		u = User{Email: email}
		if len(m.users) == 0 {
			u.IsAdmin = true
		}
		m.users[email] = u
		return u, nil
	}
	return u, nil
}

func (m *MemoryPersist) TruncateDivisions(ctx context.Context) error {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.divisions = []Division{}
	return nil
}

func (m *MemoryPersist) GetDivisions(ctx context.Context) ([]Division, error) {
	m.mx.Lock()
	defer m.mx.Unlock()
	return m.divisions, nil
}

func (m *MemoryPersist) UpsertDivisions(ctx context.Context, ds []Division) error {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.divisions = ds
	return nil
}
