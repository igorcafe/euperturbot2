package test

import (
	"context"
	"log"
	"testing"

	"github.com/igorcafe/euperturbot2/repo"
	"github.com/igorcafe/euperturbot2/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Integration_ShouldSubscribeToTopic(t *testing.T) {
	ctx := context.Background()
	log.SetFlags(log.Lshortfile)

	svc := service.Euperturbot{
		Repo: newTestRepo(t, ctx),
	}

	err := svc.SubscribeToTopic(ctx, service.SubscribeToTopicParams{
		TopicName:     "music",
		ChatID:        222,
		UserID:        111,
		UserFirstName: "Name!",
		UserUsername:  "user",
	})
	require.NoError(t, err)

	topics, err := svc.FindUserTopics(ctx, service.FindUserTopicsParams{
		ChatID: 222,
		UserID: 111,
	})
	require.NoError(t, err)

	assert.Equal(t, []string{"music"}, topics)
}

func newTestRepo(t *testing.T, ctx context.Context) repo.Repo {
	t.Helper()

	repo, err := repo.NewSQLite(ctx, ":memory:", "../repo/sqliterepo/migrations")
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, repo.Close())
	})

	return repo
}
