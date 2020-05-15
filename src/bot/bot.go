package bot

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"time"

	"github.com/keisukefd/TBot/src/Commands"
	. "github.com/keisukefd/TBot/src/Types"
)

func Run(token string, authorizedChannels []ChatID) {
	bot := configureBot(token, authorizedChannels)

	// Commands
	helpCommand := NewCustomCommand(bot, "/help", "Shows commands you can do.", Commands.Help)
	helpCommand.Invoke()

	trailerCommand := NewCustomCommand(bot, "/trailer", "Search trailer for the specified movie", Commands.Trailer)
	trailerCommand.Invoke()

	channelIDCommand := NewCustomCommand(bot, "/channelid", "Get the current channel ID", Commands.ChannelID)
	channelIDCommand.SetAdmin(true)
	channelIDCommand.SetHideHelp(true)
	channelIDCommand.Invoke()

	// Other text to parse and detect behavior
	ParsingMessage(bot)

	// Catch all callbacks from buttons
	ButtonsCallback(bot)

	bot.Start()
}

func configureBot(token string, authorizedChannels []ChatID) *tb.Bot {
	poller := &tb.LongPoller{Timeout: 15 * time.Second}
	roomRestricted := tb.NewMiddlewarePoller(poller, func(upd *tb.Update) bool {
		if upd.Message == nil {
			return true
		}
		// Autorize only a few channels ID
		if !ChatID(upd.Message.Chat.ID).In(authorizedChannels) {
			return false
		}
		return true
	})

	bot, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: roomRestricted,
	})
	if err != nil {
		log.Panic(err)
	}
	// Reset commands list
	_ = bot.SetCommands([]tb.Command{})
	return bot
}
