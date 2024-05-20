package repo

import "context"

type Repo interface {
	Close() error

	// WithTransaction(func(ctx context.Context, tx Repo) error) error

	User() UserRepo
}

type User struct {
	ID        int64  `db:"id"`
	FirstName string `db:"first_name"`
	Username  string `db:"username"`
}

type UserRepo interface {
	UserRepo()
	Save(ctx context.Context, u *User) error
	Find(ctx context.Context, u *User) error
}
