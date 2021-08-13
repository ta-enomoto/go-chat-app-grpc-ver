package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

type ConfDbUsr struct {
	Host    string `toml:"host"`
	Port    int    `toml:"port"`
	DbName  string `toml:"db-name"`
	Charset string `toml:"charset"`
	User    string `toml:"user"`
	Pass    string //環境変数の値を取得
}

type ConfigUsr struct {
	ConfDbUsr `toml:"database1"`
}

func ReadConfDbUsr() (*ConfigUsr, error) {

	confpath := "config/db.toml"

	conf := new(ConfigUsr)

	_, err := toml.DecodeFile(confpath, &conf)
	if err != nil {
		return conf, err
	}
	conf.ConfDbUsr.Pass = os.Getenv("PASSWORD1")
	return conf, nil
}
