package instance

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func openDb(conf *Config) (*gorm.DB, error) {
	var err error
	var db *gorm.DB
	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Asia%%2FTokyo",
		conf.Db.User,
		conf.Db.Password,
		conf.Db.Host,
		conf.Db.Port,
		conf.Db.Database)
	db, err = gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err == nil {
		return db, nil
	}
	return nil, err
}
