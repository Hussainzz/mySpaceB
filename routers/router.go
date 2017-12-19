package routers

import (
	"mySpaceB/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api", &controllers.MainController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login",&controllers.LoginController{})
	beego.Router("/logout",&controllers.LogoutUserController{})
	beego.Router("/newpost/",&controllers.PostController{},"get,post:Add")
	beego.Router("/post/:id([0-9]+)", &controllers.PostViewController{},"get,post:ShowP")
	beego.Router("/post/edit/:id([0-9]+)", &controllers.PostEditController{}, `get:PostEdit;post:Edit`)
	beego.Router("/post/delete/:id([0-9]+)", &controllers.PostDeleteController{}, `get:Delete`)
}
