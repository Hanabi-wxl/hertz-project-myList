package conf

import (
	"github.com/BurntSushi/toml"
	"hertz-mylist/base/logging"
	"hertz-mylist/biz/model/gorm"
	"os"
	"strings"
)

var Conf = &Config{}

type Config struct {
	Mysql *MysqlConfig
}

type MysqlConfig struct {
	DB         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
}

func InitService() {
	path, _ := os.Getwd()
	path = path + "/conf/conf.toml"
	_, err := toml.DecodeFile(path, &Conf)
	if err != nil {
		logging.Fatal("读取文件错误", err)
	}
	initMysql(Conf)
}

func initMysql(conf *Config) {
	user := conf.Mysql.DbUser
	passWord := conf.Mysql.DbPassWord
	host := conf.Mysql.DbHost
	port := conf.Mysql.DbPort
	dbName := conf.Mysql.DbName
	dsn := strings.Join([]string{user, ":", passWord, "@tcp(", host, ":", port, ")/", dbName, "?charset=utf8&parseTime=true&loc=Local"}, "")
	gorm.DateBase(dsn)
}
