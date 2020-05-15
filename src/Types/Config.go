package Types

type YamlConfig struct {
	TelegramToken      string   `yaml:"telegram_token"`
	AuthorizedChannels []string `yaml:"authorized_channels"`
}
