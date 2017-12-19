package controllers


import (
	"github.com/astaxie/beego"
	"fmt"
	models "mySpaceB/models"
	"strconv"
	"time"
)

type RegisterController struct{
	beego.Controller
}

func (r *RegisterController) Get(){
	r.TplName = "register.tpl"
}

//Registering a Users
func (r *RegisterController) Post(){
	id := strconv.FormatInt(time.Now().Unix(), 10)
	inputs := r.Input()
	uid := id
	username := inputs.Get("username")
	Pwd := inputs.Get("pwd")
	email := inputs.Get("email")
	err := models.AddUser(uid,username, Pwd, email)
	if err == nil{
		r.Redirect("/", 302)
	}

	flash := beego.ReadFromRequest(&r.Controller)
	if ok := flash.Data["error"]; ok != ""{
		r.Data["errors"] = ok
	}

	if ok := flash.Data["notice"]; ok != ""{
		r.Data["notices"] = ok
	}



}

type LoginController struct{
	beego.Controller
}

func (l *LoginController) Get(){
	l.TplName = "login.tpl"

}

func (l *LoginController) Post(){
	var user models.Users
	inputs := l.Input()
	user.Username = inputs.Get("username")
	password := inputs.Get("password")
	
	err := models.ValidateUser(user, password)
	if err == nil{
		u, _,_ := models.FindUser(user.Username, user.Id)
		
		l.SetSession("username", u) 
		l.SetSession("uname", u.Username)
		sess := l.GetSession("username")
		uname := l.GetSession("uname")

		if sess != nil {
			l.Data["InSession"] = 1
		
			l.Data["Username"] = uname
			l.Redirect("/", 302)
	}else{
		l.TplName = "register.tpl"
		fmt.Println(err)
	}
	}else{
		l.TplName="login.tpl"
		fmt.Println(err)
		
	}
}
type LogoutUserController struct {
	beego.Controller
}

func (index *LogoutUserController) Get() {
	index.DelSession("uname")
	index.DelSession("username")
	index.Redirect("/", 302)
}
