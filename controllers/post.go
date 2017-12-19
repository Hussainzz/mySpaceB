package controllers

import(
	//"github.com/astaxie/beego/validation"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
	"strconv"
	"time"
	//"io/ioutil"
	models "mySpaceB/models"
//	"encoding/json"
)

type PostController struct{
	beego.Controller
}

type PostViewController struct{
	beego.Controller
}
func (controller *PostViewController) ShowP() {
	if controller.GetSession("username") != nil {
		uname := controller.GetSession("uname")
		controller.Data["InSession"] = 1
		controller.Data["Username"] = uname
	id, _ := strconv.Atoi(controller.Ctx.Input.Param(":id"))
	post := &models.Post{Id:id}
	post.ReadDB()
	post.Users.Read()

	if post.Ispresent {
		post.Description = "POst Was Deleted"
	}

	post.Display()
	controller.Data["post"] = post
	controller.TplName = "dpost.tpl"
 }else{
	id, _ := strconv.Atoi(controller.Ctx.Input.Param(":id"))
	post := &models.Post{Id:id}
	post.ReadDB()
	post.Users.Read()

	if post.Ispresent {
		post.Description = "POst Was Deleted"
	}

	post.Display()
	controller.Data["json"] = post
	controller.ServeJSON()
 }	
}

func (this * PostController) Add(){
	
	
 
	
	if this.GetSession("username") != nil {
		users := this.GetSession("username").(models.Users)
		uname := this.GetSession("uname")
		this.Data["InSession"] = 1
		this.Data["Username"] = uname
	
		this.TplName = "newpost.tpl"
		
		o := orm.NewOrm()
		o.Using("default")
	
		if this.Ctx.Input.Method() == "POST"{
		inputs := this.Input()
		

		this.Data["Username"] = users
		//post.Author = sess.(string)
		fmt.Println("USEEEEEEEEEEEEEE", users)
		//uid, ok := se.(int)
		//fmt.Println(reflect.TypeOf(sess))
		//ii ,_:= strconv.Atoi(se.(string))

	//	fmt.Println("sdfsdfsdf----->",reflect.TypeOf(ii))
		//fmt.Println("valueeeeeee after",ii)
		post := models.Post{}
		post.Users = &users
		post.Author = uname.(string)
		post.Title = inputs.Get("title")
		post.Description = inputs.Get("description")
		post.Date = time.Now()
		

		id, err := o.Insert(&post)
		if err == nil{
			this.Redirect("/", 302)
			msg := fmt.Sprintf("POst inserted with id: %d", id)
			beego.Debug(msg)
			
		}else{
			msg := fmt.Sprintf("Couldn't insert new POst. Reason: %s", err)
			beego.Debug(msg)
		}

		}
	}else{
		this.TplName="login.tpl"
	}
}

type PostEditController struct{
	beego.Controller
}

func (controller *PostEditController) PostEdit() {
	if controller.GetSession("username") != nil {
		uname := controller.GetSession("uname")
		controller.Data["InSession"] = 1
		controller.Data["Username"] = uname
	id, _ := strconv.Atoi(controller.Ctx.Input.Param(":id"))
	post := &models.Post{Id:id}
	post.ReadDB()
	post.Users.Read()
	controller.Data["post"] = post
	controller.TplName = "edit.tpl"
	}else{
		controller.Redirect("/",320)
	}
}

func (controller *PostEditController) Edit() {
	if controller.GetSession("username") != nil {

		uname := controller.GetSession("uname")
		controller.Data["InSession"] = 1
		controller.Data["Username"] = uname

		user := controller.GetSession("username").(models.Users)
		id, _ := strconv.Atoi(controller.Ctx.Input.Param(":id"))
		post := &models.Post{Id:id}
		post.ReadDB()
		if user.Id == post.Users.Id && controller.Ctx.Input.Method() == "POST"{
			post.Title = controller.GetString("title")
			post.Description = controller.GetString("description")
		
			post.Update()
			controller.TplName = "success.tpl"
		}else{
		
		controller.Redirect("/login",320)
		}
	}else{

	}

}


type PostDeleteController struct{
	beego.Controller
}

func (d *PostDeleteController) Delete(){

	if d.GetSession("username") != nil{

		user := d.GetSession("username").(models.Users)
		
			id, _ := strconv.Atoi(d.Ctx.Input.Param(":id"))
			post := &models.Post{Id:id}
			post.ReadDB()
			if user.Id != post.Users.Id {
				
			}
		
			post.Ispresent = true
			post.Update()
			from := d.GetString("from")
			
				switch from {
				case "index":
					d.Redirect("/", 302)
				}
	}else{

	}
}