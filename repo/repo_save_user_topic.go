package repo

import (
	"context"
	"log"

	"github.com/igorcafe/euperturbot2/domain"
)

type SaveUserTopicParams struct {
	ChatID        int64
	UserID        int64
	UserFirstName string
	UserUsername  string
	TopicName     string
}

func (s Repo) SaveUserTopic(ctx context.Context, params SaveUserTopicParams) error {
	_, err := s.db.ExecContext(
		ctx,
		`INSERT INTO chat (id) VALUES ($1)
		ON CONFLICT DO NOTHING`,
		params.ChatID,
	)
	if err != nil {
		log.Print(err)
		return domain.ErrInternal
	}

	_, err = s.db.ExecContext(
		ctx,
		`INSERT INTO user
		(id, first_name, username)
		VALUES
		($1, $2, $3)
		ON CONFLICT DO NOTHING`,
		params.UserID,
		params.UserFirstName,
		params.UserUsername,
	)
	if err != nil {
		log.Print(err)
		return domain.ErrInternal
	}

	_, err = s.db.ExecContext(
		ctx,
		`INSERT INTO user_topic
		(user_id, chat_id, topic_name)
		VALUES
		($1, $2, $3)
		ON CONFLICT DO NOTHING`,
		params.UserID,
		params.ChatID,
		params.TopicName,
	)
	if err != nil {
		log.Print(err)
		return domain.ErrInternal
	}

	return nil
}
