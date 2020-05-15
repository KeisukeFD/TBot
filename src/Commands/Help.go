package Commands

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
)

func Help(bot *tb.Bot, m *tb.Message) {
	cmds, _ := bot.GetCommands()

	message := "Here's commands available: \n"
	for _, c := range cmds {
		message += fmt.Sprintf("- /%-15s %s\n", c.Text, c.Description)
	}
	bot.Send(m.Chat, message)
}
