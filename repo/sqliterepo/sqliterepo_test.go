package sqliterepo_test

import (
	"context"
	"testing"

	"github.com/igorcafe/euperturbot2/repo"
	"github.com/igorcafe/euperturbot2/repo/sqliterepo"
	"github.com/stretchr/testify/require"
)

func newTestRepo(t *testing.T, ctx context.Context) repo.Repo {
	t.Helper()

	repo, err := sqliterepo.NewSQLite(ctx, ":memory:", "./migrations")
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, repo.Close())
	})

	return repo
}
