package Types

type ChatID int64

func (i ChatID) In(list []ChatID) bool {
	for _, j := range list {
		if i == j {
			return true
		}
	}
	return false
}
