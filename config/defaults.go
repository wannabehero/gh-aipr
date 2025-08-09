package config

import "github.com/spf13/viper"

func setDefaults() {
	viper.SetDefault("anthropic.model_name", "claude-3-5-haiku-latest")
	viper.SetDefault("gemini.model_name", "gemini-2.5-flash")
	viper.SetDefault("openai.model_name", "gpt-4o-mini")
}
