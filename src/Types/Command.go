package Types

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

type CustomCommand struct {
	Cmd    tb.Command
	Bot    *tb.Bot
	Run    func(*tb.Bot, *tb.Message)
	Hidden bool
	Admin  bool
}

func NewCustomCommand(bot *tb.Bot, handle string, description string, function func(*tb.Bot, *tb.Message)) CustomCommand {
	return CustomCommand{
		Cmd: tb.Command{
			Text:        handle,
			Description: description,
		},
		Bot:    bot,
		Run:    function,
		Hidden: false,
		Admin:  false,
	}
}

func (c CustomCommand) Invoke() {
	c.Bot.Handle(c.Cmd.Text, func(m *tb.Message) {
		if c.Admin {
			adminMembers, err := c.Bot.AdminsOf(m.Chat)
			if err != nil {
				return
			}
			for _, member := range adminMembers {
				if m.Sender.Username == member.User.Username {
					c.Run(c.Bot, m)
					break
				}
			}
		} else {
			c.Run(c.Bot, m)
		}
	})
	if !c.Hidden {
		setCommand(c.Bot, c.Cmd)
	}
}

// Hide the command from Help (ex. for admin purposes)
func (c *CustomCommand) SetHideHelp(value bool) {
	c.Hidden = value
}

// Must be Admin (or creator) to run the command
func (c *CustomCommand) SetAdmin(value bool) {
	c.Admin = value
}

func setCommand(bot *tb.Bot, cmd tb.Command) {
	commands, _ := bot.GetCommands()
	var newCommands []tb.Command
	for _, c := range commands {
		newCommands = append(newCommands, c)
	}
	newCommands = append(newCommands, cmd)
	_ = bot.SetCommands(newCommands)
}
