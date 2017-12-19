package models

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
	"github.com/gogather/com"
)

type Users struct{
	Id string `orm:"pk" json:"id"`
	Username string `json:"username"`
	Password string	`json:"pwd"`
	Salt string	`json:"salt"`
	Email string `orm:"unique" json:"email"`
}

func (this *Users) TableName() string{
	return "users"
}

type Post struct{
	Id int `orm:"pk;auto" json:"pid"`
	Users *Users `orm:"rel(fk)" json:"users"`
	Author string `json:"author"`
	Title string `form:"title,text,Title:" valid:"MinSize(5); MaxSize(20)" json:"title"`
	Description string `orm:"size(12000)" json:"desc" `
	Date time.Time `orm:"auto_now_add;type(datetime)" json:"date"`
	Ispresent bool `json:"isprsnt"`
}



func (post *Post) Display(){
	fmt.Printf("Post Id : %s\n", post.Id)
	fmt.Printf("User Id(foreignKey) : %s\n", post.Users)
	fmt.Printf("Author  : %s\n", post.Author)
	fmt.Printf("Title  : %s\n", post.Title)
	fmt.Printf("Description : %s\n", post.Description)
	fmt.Printf("Date : %s\n",post.Date)
	fmt.Printf("Ispresent: %s\n", post.Ispresent)
}



func init(){
	orm.RegisterModel(new(Users), new(Post))
}

func AddUser(uid string,username string, password string, email string)(error){
	o := orm.NewOrm()
	o.Using("default")
	user := new(Users)
	user.Id = uid
	user.Username = username
	user.Salt = com.RandString(10)
	user.Password = com.Md5(password + user.Salt)
	user.Email = email
	o.Insert(user)
	return nil
}

func FindUser(username string, id string) (Users ,error, error){
	o := orm.NewOrm()
	o.Using("default")
	user := Users{Username:username, Id:id}
	err := o.Read(&user,"username")
	errr := o.Read(&user, "id")
	return user, err, errr
}

func ValidateUser(user Users, password string) error{
	flash := beego.NewFlash()
	u, _ , _ := FindUser(user.Username, user.Id)
	if u.Username == "" {
		flash.Error("wrong user name or password!")
		return errors.New("wrong user name or password!")
	}
	fmt.Println(u.Id)
	fmt.Println(u.Password)
	fmt.Println(com.Md5(password + u.Salt))
	if u.Password != com.Md5(password + u.Salt) {
		flash.Error("Wrong Password TryAgain")
		
		return errors.New("wrong password!")
	}
	return nil
}



func (post *Post) ReadDB() (err error) {
	o := orm.NewOrm()
	if err = o.Read(post); err != nil {
		beego.Info(err)
	}

	return
}

func (users *Users)Read() (err error) {
	o := orm.NewOrm()
	err = o.Read(users)
	return
}

func (post Post) GetAll() (ret []Post) {
	o := orm.NewOrm()
	o.QueryTable("post").OrderBy("-date").Filter("ispresent", 0).All(&ret)
	return
}

func (post Post) Gets() (ret []Post) {
	o := orm.NewOrm()
	o.QueryTable("post").Filter("Users", post.Users).Filter("ispresent", 0).All(&ret)
	return
}

func (post Post) Update() (err error) {
	o := orm.NewOrm()
	if _, err = o.Update(&post); err != nil {
		beego.Info(err)
		
	}
	return
}