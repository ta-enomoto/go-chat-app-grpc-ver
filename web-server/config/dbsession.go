package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

type ConfDbSession struct {
	Host    string `toml:"host"`
	Port    int    `toml:"port"`
	DbName  string `toml:"db-name"`
	Charset string `toml:"charset"`
	User    string `toml:"user"`
	Pass    string //環境変数の値を取得
}

type ConfigSession struct {
	ConfDbSession `toml:"database1"`
}

func ReadConfDbSession() (*ConfigSession, error) {

	confpath := "config/db.toml"

	conf := new(ConfigSession)

	_, err := toml.DecodeFile(confpath, &conf)
	if err != nil {
		return conf, err
	}
	conf.ConfDbSession.Pass = os.Getenv("PASSWORD1")
	return conf, nil
}
