package bot_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/igorcafe/euperturbot2/bot"
)

func TestHandlesWebhook(t *testing.T) {
	// Arrange
	myBot, err := bot.New(bot.NewParams{
		Token: "ma token",
	})
	if err != nil {
		t.Fatal(err)
	}

	update := bot.Update{
		Message: &bot.Message{
			Chat: &bot.Chat{
				ID: 123,
			},
			Text: "some text",
		},
	}
	bWant, _ := json.Marshal(update)
	req := httptest.NewRequest("GET", "/doesnt/matter", bytes.NewReader(bWant))
	rec := httptest.NewRecorder()

	// Act
	myBot.HandleWebhook(rec, req)

	// Check
	gotUpdate := <-myBot.Updates()

	if rec.Result().StatusCode != http.StatusOK {
		t.Fatal(http.StatusText(rec.Result().StatusCode))
	}

	bGot, _ := json.Marshal(gotUpdate)

	if !bytes.Equal(bWant, bGot) {
		t.Fatalf("updates - want: %s, got: %s", string(bWant), string(bGot))
	}

	t.Log(string(bGot))
}
