package main

import (
	"github.com/astaxie/beego/orm"
	_ "mySpaceB/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	_ "mySpaceB/routers"
	_ "mySpaceB/models"
	//"github.com/astaxie/beego/context"
)

func init(){
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default",  "mysql",  "root:admin@/myspace?charset=utf8")
	name := "default"
	force := false
	verbose := false
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil{
		fmt.Println(err)
	}

}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	
	beego.Run()
}



