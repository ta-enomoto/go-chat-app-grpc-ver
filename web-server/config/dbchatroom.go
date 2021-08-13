package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

type ConfDbChtrm struct {
	Host    string `toml:"host"`
	Port    int    `toml:"port"`
	DbName  string `toml:"db-name"`
	Charset string `toml:"charset"`
	User    string `toml:"user"`
	Pass    string //環境変数の値を取得
}

type ConfigChtrm struct {
	ConfDbChtrm `toml:"database2"`
}

func ReadConfDbChtrm() (*ConfigChtrm, error) {

	confpath := "config/db.toml"

	conf := new(ConfigChtrm)

	_, err := toml.DecodeFile(confpath, &conf)
	if err != nil {
		return conf, err
	}
	conf.ConfDbChtrm.Pass = os.Getenv("PASSWORD1")
	return conf, nil
}
