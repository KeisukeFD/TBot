package bot

import (
	"github.com/keisukefd/TBot/src/Commands"
	"github.com/keisukefd/TBot/src/Types"
	"github.com/keisukefd/TBot/src/filters"
	tb "gopkg.in/tucnak/telebot.v2"
	"strings"
)

func ParsingMessage(bot *tb.Bot) {
	bot.Handle(tb.OnText, func(m *tb.Message) {
		msg := Types.Message{
			Content: m,
			Bot:     bot,
		}

		msg.Filters(
			filters.DetectHello,
			filters.DetectMagic,
		)
	})
}

func ButtonsCallback(bot *tb.Bot) {
	bot.Handle(tb.OnCallback, func(c *tb.Callback) {
		actionName := strings.Split(strings.TrimSpace(c.Data), ":")
		directResponse := "👍"
		switch actionName[0] {
		case Commands.TrailerURI:
			directResponse = "🔎"
			resp := Commands.TrailerOnCallback(c)
			bot.Send(c.Message.Chat, resp)
		default:
			return
		}

		bot.Respond(c, &tb.CallbackResponse{
			Text: directResponse,
		})
	})
}
