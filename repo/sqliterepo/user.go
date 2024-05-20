package sqliterepo

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/igorcafe/euperturbot2/domain"
	"github.com/igorcafe/euperturbot2/repo"
	"github.com/jmoiron/sqlx"
)

var _ repo.UserRepo = userRepo{}

type userRepo struct {
	// db executor
	db *sqlx.DB
}

func (userRepo) UserRepo() {}

func (r userRepo) Save(ctx context.Context, u *repo.User) error {
	_, err := r.db.NamedExecContext(ctx, `
	INSERT INTO user
	(
		id,
		first_name,
		username
	)
	VALUES (
		:id,
		:first_name,
		:username
	)
	ON CONFLICT DO UPDATE
	SET
		id = :id,
		first_name = :first_name,
		username = :username
	`, u)
	if err != nil {
		log.Print("ERR: ", err)
		return domain.ErrInternal
	}
	return nil
}

func (r userRepo) Find(ctx context.Context, u *repo.User) error {
	err := r.db.QueryRowxContext(ctx, `
	SELECT
		id,
		first_name,
		username
	FROM
		user
	WHERE
		id = $1
	`, u.ID).StructScan(u)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.ErrUserNotFound
	}
	if err != nil {
		log.Print("ERR: ", err)
		return domain.ErrInternal
	}
	return nil
}
