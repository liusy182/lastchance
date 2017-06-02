package models

type User struct {
	Name string `valid:"Required;MaxSize(15)"`
	Age  int    `valid:"Range(0,18)"`
}
