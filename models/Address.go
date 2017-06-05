package models

type Address struct {
	ID   int   // ID would create field "i_d" in database
	User *User `orm:"reverse(one)"`
}
