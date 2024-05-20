package sqliterepo_test

import (
	"context"
	"testing"

	"github.com/igorcafe/euperturbot2/repo"
	"github.com/stretchr/testify/require"
)

func Test_User_ShouldRetrieveSavedUser(t *testing.T) {
	// Arrange
	ctx := context.Background()
	myrepo := newTestRepo(t, ctx)

	wantUser := repo.User{
		ID:        123,
		FirstName: "José",
		Username:  "lemosbash",
	}

	// Act
	err := myrepo.User().Save(ctx, &wantUser)
	require.NoError(t, err)

	gotUser := repo.User{ID: 123}
	err = myrepo.User().Find(ctx, &gotUser)
	require.NoError(t, err)

	// Check
	require.Equal(t, wantUser, gotUser)
}
