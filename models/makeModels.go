package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var x *xorm.Engine
var err error

func init() {
	user_name := beego.AppConfig.String("myg_user")
	user_pwd := beego.AppConfig.String("myg_passwd")
	host := beego.AppConfig.String("myg_ip")
	port := beego.AppConfig.String("myg_port")
	db := beego.AppConfig.String("myg_db")
	x, err = xorm.NewEngine("mysql", user_name+":"+user_pwd+"@tcp("+host+":"+port+")/"+db+"?charset=utf8")
	if err != nil {
		logs.Error("数据库连接失败")
	}
}

func GetDbTables() []map[string][]byte {
	db := beego.AppConfig.String("myg_db")
	list, err := x.Query("select TABLE_NAME from information_schema.`TABLES` WHERE TABLE_SCHEMA  = '" + db + "'")
	if err != nil {
		logs.Error("获取数据库数据表失败")
	}
	return list
}

func GetTableColum(table_name string) []map[string][]byte {
	db := beego.AppConfig.String("myg_db")
	list, err := x.Query("select COLUMN_NAME,DATA_TYPE,COLUMN_TYPE,COLUMN_KEY,COLUMN_COMMENT from information_schema.`COLUMNS` WHERE TABLE_SCHEMA  = '" + db + "' AND TABLE_NAME = '" + table_name + "'")
	if err != nil {
		logs.Error("获取数据表结构失败")
	}
	return list
}
