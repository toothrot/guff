package models

import (
	"context"
	"database/sql"
	"errors"
	"sync"
)

type DBPersist struct {
	DB *sql.DB
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

var DefaultMemoryPersist = &MemoryPersist{divisions: []Division{}, mx: &sync.Mutex{}}

type MemoryPersist struct {
	divisions []Division
	mx        *sync.Mutex
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

type Persist interface {
	UpsertDivisions(ctx context.Context, ds []Division) error
	GetDivisions(ctx context.Context) ([]Division, error)
	TruncateDivisions(ctx context.Context) error
}
