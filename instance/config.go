package instance

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Db struct {
		Host     string `yaml:"host"`     // MySQLの接続先ホスト
		Port     int    `yaml:"port"`     // MySQLの接続先ポート
		User     string `yaml:"user"`     // MySQLの接続先ユーザー名
		Password string `yaml:"password"` // MySQLの接続先パスワード
		Database string `yaml:"database"` // MySQLの接続先データベース
	} `yaml:"db"`
}

func NewConfig() *Config {
	return &Config{}
}

func (a *Config) Load(configFile string) error {
	buf, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(buf, a)
	if err != nil {
		return err
	}

	return nil
}
