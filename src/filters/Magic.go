package filters

import (
	"github.com/keisukefd/TBot/src/Types"
	"regexp"
)

func DetectMagic(m Types.Message) bool {
	var validString = regexp.MustCompile(`(?i)(magic)`)
	if validString.MatchString(m.Content.Text) {
		m.Bot.Send(m.Content.Sender, "Where Magic happens !")
		return true
	}
	return false
}
