package instance

import (
	"gorm.io/gorm"
)

type Instance struct {
	Db   *gorm.DB
	Conf *Config
}

func NewInstance() *Instance {
	return &Instance{}
}

func (a *Instance) Init(configPath, firebasePath string) error {
	var err error
	conf := NewConfig()
	err = conf.Load(configPath)
	if err != nil {
		return err
	}
	a.Db, err = openDb(conf)
	if err != nil {
		return err
	}
	a.Conf = conf

	return nil
}
