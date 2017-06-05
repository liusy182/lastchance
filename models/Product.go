package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Product struct {
	ID           int       `orm:"pk;auto"`
	Name         string    `orm:"index"`
	Description  string    `orm:"column(product_desc)"`
	SerialNumber int       `orm:"size(15)"`
	Value        float32   `orm:"digits(10);decimals(2)"`
	Inventory    int       `orm:"-"`
	LastOrdered  time.Time `orm:"auto_now_add;type(date)"`
}

func init() {
	orm.RegisterModel(new(Product))
}

func (p *Product) TableIndex() [][]string {
	return [][]string{
		[]string{"SerialNumber", "LastOrdered"},
	}
}
