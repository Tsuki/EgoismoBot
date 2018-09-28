package config

type Config struct {
	Debug         bool `toml:"debug" json:"debug"`
	LoggingConfig `toml:"logging"`
}

type LoggingConfig struct {
	Level        string `toml:"level" json:"level"`
	Output       string `toml:"output" json:"output"`
	TelegramHook string `toml:"telegram_hook" json:"telegram_hook"`
	AppName      string `toml:"app_name" json:"app_name"`
}
