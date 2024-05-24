package bot

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type Bot interface {
	HandleWebhook(w http.ResponseWriter, r *http.Request)
	Updates() <-chan Update
}

type bot struct {
	token   string
	updates chan Update
}

type NewParams struct {
	Token string
}

func New(params NewParams) (Bot, error) {
	return bot{
		token:   params.Token,
		updates: make(chan Update, 1),
	}, nil
}

func (b bot) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	var u Update
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Print("failed parse update: ", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	b.updates <- u
}

func (b bot) Updates() <-chan Update {
	return b.updates
}

func (b bot) SendMessage(ctx context.Context, params any) error {
	return nil
}

type Update struct {
	Message *Message `json:"message,omitempty"`
}

type Message struct {
	Chat *Chat  `json:"chat,omitempty"`
	From *User  `json:"from,omitempty"`
	Text string `json:"text,omitempty"`
}

type User struct {
	ID        int64  `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	FirstName string `json:"first_name,omitempty"`
}

type Chat struct {
	ID int64 `json:"id,omitempty"`
}
