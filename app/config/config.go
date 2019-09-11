package config

import (
	"bytes"
	"github.com/eddieowens/axon"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"log"
	"path"
	"runtime"
	"strings"
)

const Key = "Config"

type Config struct {
	Git    Git    `mapstructure:"git"`
	Server Server `mapstructure:"server"`
	Log    Log    `mapstructure:"log"`
}

type Log struct {
	Level      string `mapstructure:"level"`
	TimeFormat string `mapstructure:"timeformat"`
}

type Git struct {
	Remote          string `mapstructure:"remote"`
	Branch          string `mapstructure:"branch"`
	Directory       string `mapstructure:"directory"`
	PollingInterval int    `mapstructure:"pollinginterval"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	SSHKey          string `mapstructure:"sshkey"`
}

type Server struct {
	Port int `mapstructure:"port"`
}

func defaultConfig() *Config {
	return &Config{}
}

func configFactory(_ axon.Injector, _ axon.Args) axon.Instance {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetEnvPrefix("simaia")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	_, filename, _, _ := runtime.Caller(0)
	d, _ := path.Split(filename)
	v.AddConfigPath(v.GetString("CONFIG_DIR"))
	v.AddConfigPath(path.Join(d, "..", "..", "config"))
	v.AddConfigPath("./config")

	b, _ := yaml.Marshal(defaultConfig())
	defaultConfig := bytes.NewReader(b)
	if err := v.MergeConfig(defaultConfig); err != nil {
		panic(err)
	}

	v.SetConfigName("config")
	if err := v.MergeInConfig(); err != nil {
		panic(err)
	}

	v.AutomaticEnv()

	config := Config{}
	if err := v.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return axon.Any(config)
}
