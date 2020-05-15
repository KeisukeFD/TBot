package Commands

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
)

func ChannelID(bot *tb.Bot, m *tb.Message) {
	message := fmt.Sprintf("The Channel ID for '%s' is '%d'", m.Chat.Title, m.Chat.ID)
	bot.Send(m.Sender, message)
	bot.Delete(m)
}
