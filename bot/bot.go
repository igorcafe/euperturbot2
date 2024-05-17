package bot

import (
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

type Update struct {
	Message *Message `json:"message,omitempty"`
}

type Message struct {
	Chat *Chat  `json:"chat,omitempty"`
	Text string `json:"text,omitempty"`
}

type Chat struct {
	ID int64 `json:"id,omitempty"`
}
