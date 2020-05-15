package Types

import tb "gopkg.in/tucnak/telebot.v2"

type Message struct {
	Content *tb.Message
	Bot     *tb.Bot
}

func (m Message) Filters(funcs ...func(m Message) bool) {
	for _, f := range funcs {
		if f(m) {
			break
		}
	}
}
