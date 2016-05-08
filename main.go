package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/favframework/debug"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"path/filepath"
	_ "photolimit/models"
	_ "photolimit/routers"
)

func init() {
	basedir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	beego.AppConfig.Set("basedir", basedir)
	cc := beego.AppConfig.String("basedir")
	godump.Dump(cc)

	orm.RegisterDataBase("default", "mysql", "root:@/limitpic?charset=utf8")
	orm.RunSyncdb("default", true, true)
}

func main() {

	// beego.BConfig.Log.AccessLogs = true
	orm.Debug = true
	beego.Run()
	// orm.RunCommand()
}
