package config

import (
	"gopkg.in/ini.v1"
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
}
