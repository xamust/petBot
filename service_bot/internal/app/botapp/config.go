package botapp

type Config struct {
	APIKey     string `toml:"api_key"`
	OwnerId    int    `toml:"owner_id"`
	LogLevel   string `toml:"log_level"`
	BotDebug   bool   `toml:"bot_debug"`
	BotTimeout int    `toml:"bot_timeout"`
}

func NewConfig() *Config {
	return &Config{
		LogLevel:   "info",
		BotDebug:   false,
		BotTimeout: 60,
	}
}
