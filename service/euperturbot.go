package service

import (
	"context"

	"github.com/igorcafe/euperturbot2/domain"
	"github.com/igorcafe/euperturbot2/repo"
)

type Euperturbot struct {
	Repo repo.Repo
}

type SubscribeToTopicParams struct {
	TopicName     string
	ChatID        int64
	UserID        int64
	UserFirstName string
	UserUsername  string
}

func (s Euperturbot) SubscribeToTopic(ctx context.Context, params SubscribeToTopicParams) error {
	err := s.Repo.SaveUserTopic(ctx, repo.SaveUserTopicParams{
		ChatID:        params.ChatID,
		UserID:        params.UserID,
		TopicName:     params.TopicName,
		UserFirstName: params.UserFirstName,
		UserUsername:  params.UserUsername,
	})
	if err != nil {
		return domain.ErrInternal
	}

	return nil
}

type FindUserTopicsParams struct {
	ChatID int64
	UserID int64
}

func (s Euperturbot) FindUserTopics(ctx context.Context, params FindUserTopicsParams) ([]string, error) {
	return s.Repo.FindUserTopics(ctx, repo.FindUserTopicParams{
		UserID: params.UserID,
		ChatID: params.ChatID,
	})
}
