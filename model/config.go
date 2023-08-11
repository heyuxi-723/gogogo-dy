package model

type Config struct {
	Port  string
	MySql Mysql
}

type Mysql struct {
	DbUser    string `json:db_user`
	DbName    string `json:db_name`
	DbPwd     string `json:db_pwd`
	DbHost    string `json:db_host`
	DbPort    string `json:db_port`
	DbCharset string `json:db_charset`
}
