package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	TelegramToken     string    `mapstructure:"telegram_token"`
	PocketConsumerKey string    `mapstructure:"pocket_consumer_key"`
	AuthServerUrl     string    `mapstructure:"auth_server_url"`
	TelegramBotUrl    string    `mapstructure:"telegram_bot_url"`
	DbFile            string    `mapstructure:"db_file"`
	BotUrl            string    `mapstructure:"bot_url"`
	Messages          *Messages `mapstructure:"messages"`
	Commands          *Commands `mapstructure:"commands"`
}

type Messages struct {
	Errors    *Errors    `mapstructure:"errors"`
	Responses *Responses `mapstructure:"responses"`
}

type Commands struct {
	Start string `mapstructure:"start"`
}

type Responses struct {
	Start                string `mapstructure:"start"`
	AlreadyAuthorized    string `mapstructure:"already_authorized"`
	UrlSavedSuccessfully string `mapstructure:"url_saved_successfully"`
}

type Errors struct {
	InvalidUrl     string `mapstructure:"invalid_url"`
	Unauthorized   string `mapstructure:"unauthorized"`
	UnableToSave   string `mapstructure:"unable_to_save"`
	UnknownCommand string `mapstructure:"unknown_command"`
	UnknownError   string `mapstructure:"unknown_error"`
}

func Init() (*Config, error) {
	var cfg Config

	//viper.SetConfigName("configs")
	viper.SetConfigName("main")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.responses", &cfg.Messages.Responses); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.errors", &cfg.Messages.Errors); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("commands", &cfg.Messages.Errors); err != nil {
		return nil, err
	}

	if err := parseEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func parseEnv(cfg *Config) error {
	os.Setenv("TOKEN", "5516775333:AAHEepCjVdoZyLPI56WCZteps_SAYRZS_84") //залупа для примера
	os.Setenv("CONSUMER_KEY", "102608-f22e24e51debd40fdedd925")
	os.Setenv("AUTH_SERVER_URL", "http://localhost/")
	os.Setenv("TELEGRAM_BOT_URL", "http://t.me/pocket_study_for_bot")

	if err := viper.BindEnv("token"); err != nil {
		return err
	}
	if err := viper.BindEnv("consumer_key"); err != nil {
		return err
	}
	if err := viper.BindEnv("auth_server_url"); err != nil {
		return err
	}
	if err := viper.BindEnv("telegram_bot_url"); err != nil {
		return err
	}

	cfg.TelegramToken = viper.GetString("token")
	cfg.PocketConsumerKey = viper.GetString("consumer_key")
	cfg.AuthServerUrl = viper.GetString("auth_server_url")
	cfg.TelegramBotUrl = viper.GetString("telegram_bot_url")

	return nil
}
