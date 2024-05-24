package controller

import (
	"context"

	"github.com/igorcafe/euperturbot2/bot"
	"github.com/igorcafe/euperturbot2/repo"
)

func SubscribeToTopic(ctx context.Context, b bot.Bot, update bot.Update) error {
	var r repo.Repo
	err := r.SaveUserTopic(ctx, repo.SaveUserTopicParams{
		ChatID:        update.Message.Chat.ID,
		UserID:        update.Message.From.ID,
		UserFirstName: update.Message.From.FirstName,
		UserUsername:  update.Message.From.Username,
	})
	if err != nil {

	}
}
