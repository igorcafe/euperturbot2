package repo

import (
	"context"
	"log"

	"github.com/igorcafe/euperturbot2/domain"
)

type FindUserTopicParams struct {
	UserID int64
	ChatID int64
}

func (s Repo) FindUserTopics(ctx context.Context, params FindUserTopicParams) ([]string, error) {
	var res []string
	err := s.db.SelectContext(
		ctx,
		&res,
		`SELECT ut.topic_name FROM user_topic ut
		WHERE ut.user_id = $1 AND ut.chat_id = $2`,
		params.UserID,
		params.ChatID,
	)
	if err != nil {
		log.Print(err)
		return nil, domain.ErrInternal
	}
	return res, nil
}
