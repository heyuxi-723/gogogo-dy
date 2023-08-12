package config

import (
	"github.com/RaymondCode/simple-demo/models"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
)

var Cfg *ini.File
var Config Conf

type Conf struct {
	Port  string
	MySql Mysql
	Jwt   Jwt
	DB    *gorm.DB
}

type Mysql struct {
	DbUser    string `json:db_user`
	DbName    string `json:db_name`
	DbPwd     string `json:db_pwd`
	DbHost    string `json:db_host`
	DbPort    string `json:db_port`
	DbCharset string `json:db_charset`
}

type Jwt struct {
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"`
	JwtTtl int64  `mapstructure:"jwt_ttl" json:"jwt_ttl" yaml:"jwt_ttl"` // token 有效期（秒）
}

func InitConfig() {
	var err error
	// 获取当前工作目录
	currentDir, _ := os.Getwd()

	// 构建配置文件的绝对路径
	configPath := filepath.Join(currentDir, "config/config.ini")
	Cfg, err = ini.Load(configPath)
	if err != nil {
		log.Fatalf("Fail to parse 'config.ini': %v", err)
	}

	loadApp()
	loadMysql()
}

func loadApp() {
	Config.Port = Cfg.Section("app").Key("Port").String()
	Config.Jwt.Secret = Cfg.Section("jwt").Key("secret").String()
	Config.Jwt.JwtTtl, _ = Cfg.Section("jwt").Key("jwt_ttl").Int64()
}

func loadMysql() {
	Config.MySql = Mysql{
		DbName:    Cfg.Section("mysql").Key("db_name").String(),
		DbUser:    Cfg.Section("mysql").Key("db_user").String(),
		DbPwd:     Cfg.Section("mysql").Key("db_pwd").String(),
		DbHost:    Cfg.Section("mysql").Key("db_host").String(),
		DbPort:    Cfg.Section("mysql").Key("db_port").String(),
		DbCharset: Cfg.Section("mysql").Key("db_charset").String(),
	}

	dataSource := Config.MySql.DbUser + ":" + Config.MySql.DbPwd + "@tcp(" + Config.MySql.DbHost + ":" + Config.MySql.DbPort + ")/" + Config.MySql.DbName + "?charset=" + Config.MySql.DbCharset + "&parseTime=true"
	DB, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{
		PrepareStmt:            true, //缓存预编译命令
		SkipDefaultTransaction: true, //禁用默认事务操作
	})
	if err != nil {
		log.Fatalf("Fail to connect DB: %v", err)
	}

	// 数据库表动态迁移
	DB.AutoMigrate(&models.User{})
	Config.DB = DB
}
