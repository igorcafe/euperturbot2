package sqliterepo

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"

	"github.com/igorcafe/euperturbot2/repo"
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

var _ repo.Repo = sqliteRepo{}

type sqliteRepo struct {
	db   executor
	user repo.UserRepo
}

type executor interface {
	Close() error

	NamedExecContext(ctx context.Context, query string, arg any) (sql.Result, error)
	NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error)
}

func NewSQLite(ctx context.Context, path string, dir string) (repo.Repo, error) {
	db, err := sqlx.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	query, err := os.ReadFile(filepath.Join(dir, "1_create_table_user.sql"))
	if err != nil {
		return nil, err
	}

	_, err = db.ExecContext(ctx, string(query))
	if err != nil {
		return nil, err
	}

	return sqliteRepo{
		db,
		userRepo{db},
	}, nil
}

func (s sqliteRepo) Close() error {
	return s.db.Close()
}

// User implements repo.Repo.
func (s sqliteRepo) User() repo.UserRepo {
	return s.user
}
