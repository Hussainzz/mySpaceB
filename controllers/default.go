package controllers

import (
	"github.com/astaxie/beego"
	"mySpaceB/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	if c.GetSession("username") != nil{
		uname := c.GetSession("uname")
		user := c.GetSession("username").(models.Users)
		c.Data["user"] = user
		c.Data["InSession"] = 1
		c.Data["Username"] = uname
		post := &models.Post{}
		post_list := post.GetAll()
		for _, p := range post_list {
			p.Users.Read()
		}
		c.Data["post_list"] = post_list
		c.TplName = "index.tpl"
	}else{
	post := &models.Post{}
	post_list := post.GetAll()

	for _, p := range post_list {
		p.Users.Read()
	}

	c.Data["json"] = post_list
	//c.TplName = "index.tpl"

	c.ServeJSON()
	}



}
