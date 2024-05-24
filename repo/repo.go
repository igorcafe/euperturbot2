package repo

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

type Repo struct {
	db *sqlx.DB
}

func NewSQLite(ctx context.Context, path string, dir string) (Repo, error) {
	repo := Repo{}
	db, err := sqlx.Open("sqlite", path)
	if err != nil {
		return repo, err
	}

	err = migrate(ctx, db, dir)
	if err != nil {
		return repo, err
	}

	_, err = db.ExecContext(ctx, `PRAGMA foreign_keys = ON`)
	if err != nil {
		return repo, err
	}

	return Repo{
		db,
	}, nil
}

func migrate(ctx context.Context, db *sqlx.DB, dir string) error {
	migrations, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	var version int
	err = db.QueryRowContext(ctx, `PRAGMA user_version`).Scan(&version)
	if err != nil {
		return err
	}
	version++

	for {
		i := slices.IndexFunc(migrations, func(m fs.DirEntry) bool {
			return strings.HasPrefix(m.Name(), fmt.Sprintf("%d_", version))
		})

		if i == -1 {
			log.Print("INF: migration finished")
			return nil
		}

		migration := filepath.Join(dir, migrations[i].Name())
		log.Printf("INF: RUN migration %s", migration)

		err = func() error {
			var err error

			ctx, cancel := context.WithTimeout(ctx, time.Second)
			defer cancel()

			tx, err := db.BeginTx(ctx, nil)
			if err != nil {
				return err
			}
			defer func() {
				err = errors.Join(err, tx.Rollback())
			}()

			query, err := os.ReadFile(migration)
			if err != nil {
				return err
			}

			_, err = tx.ExecContext(ctx, string(query))
			if err != nil {
				return err
			}

			version++
			_, err = tx.ExecContext(ctx, fmt.Sprintf("PRAGMA user_version = %d", version))
			if err != nil {
				return err
			}

			return tx.Commit()
		}()

		if err != nil {
			log.Printf("INF: FAIL migration %s", migration)
			return err
		}
		log.Printf("INF: SUCCEED migration %s", migration)
	}
}

func (s Repo) Close() error {
	return s.db.Close()
}
