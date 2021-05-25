package main

import (
	"encoding/gob"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"myApp/models"
	_ "myApp/models"
	_ "myApp/routers"
)

func main() {
	sysInit()
	dbInit()
	logs.SetLogger("console")
	beego.Run()
}

//sysInit :系统初始化
func sysInit() {
	gob.Register(models.Member{}) //序列化Member对象,必须在encoding/gob编码解码前进行注册
}

//dbInit :数据库初始化
func dbInit(aliases ...string) {
	isDev := beego.AppConfig.String("runmode") == "dev"

	if len(aliases) > 0 {
		for _, alias := range aliases {
			registDatabase(alias)
			// 自动建表
			if "w" == alias {
				orm.RunSyncdb("default", false, isDev)
			}
		}
	} else {
		registDatabase("w")
		orm.RunSyncdb("default", false, isDev)
	}

	if isDev {
		orm.Debug = isDev
	}

}

func registDatabase(alias string) {
	if len(alias) == 0 {
		return
	}
	//连接名称
	dbAlias := alias
	if "w" == alias || "default" == alias {
		dbAlias = "default"
		alias = "w"
	}
	//数据库名称
	dbName := beego.AppConfig.String("db_" + alias + "_database")
	//数据库连接用户名
	dbUser := beego.AppConfig.String("db_" + alias + "_username")
	//数据库连接用户名
	dbPwd := beego.AppConfig.String("db_" + alias + "_password")
	//数据库IP（域名）
	dbHost := beego.AppConfig.String("db_" + alias + "_host")
	//数据库端口
	dbPort := beego.AppConfig.String("db_" + alias + "_port")

	err := orm.RegisterDataBase(dbAlias, "mysql", dbUser+":"+dbPwd+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8", 30)
	if err != nil {
		return
	}

}
