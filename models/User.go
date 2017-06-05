package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id   int    // ID would create field "i_d" in database
	Name string `valid:"Required;MaxSize(15)"`
	Age  int    `valid:"Range(0,18)"`
}

func init() {
	orm.RegisterModel(new(User))
}

//customize table name
func (u *User) TableName() string {
	return "custom_user"
}
