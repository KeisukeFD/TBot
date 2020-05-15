package filters

import (
	"fmt"
	"github.com/keisukefd/TBot/src/Types"
	"regexp"
)

func DetectHello(m Types.Message) bool {
	var validString = regexp.MustCompile(`(?i)^(hello|allo|bonjour|salut|salu|hey|yo)`)
	if validString.MatchString(m.Content.Text) {
		m.Bot.Send(m.Content.Chat, fmt.Sprintf("Salut @%s !", m.Content.Sender.Username))
		return true
	}
	return false
}
