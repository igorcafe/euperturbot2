package function

import (
	"encoding/json"
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/igorcafe/euperturbot2/bot"
)

func init() {
	// projectID := os.Getenv("PROJECT_ID")
	myBot, err := bot.New(bot.NewParams{
		Token: os.Getenv("BOT_TOKEN"),
	})
	if err != nil {
		log.Panic(err)
	}
	functions.HTTP("bot", myBot.HandleWebhook)

	go func() {
		for update := range myBot.Updates() {
			b, _ := json.Marshal(update)
			log.Print("update: ", string(b))
		}
	}()
}
