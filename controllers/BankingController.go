package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/liusy182/lastchance/models"
)

type BankingController struct {
	beego.Controller
}

func (c *BankingController) URLMapping() {
	c.Mapping("ShowAccounts", c.ShowAccounts)
	c.Mapping("Transfer", c.Transfer)
}

// @router /banking [get]
func (c *BankingController) ShowAccounts() {
	c.Data["accounts"] = []models.Account{
		models.Account{
			ID:     1,
			Name:   "Checking",
			Number: "8888",
			Amount: 45.22,
		},
		models.Account{
			ID:     2,
			Name:   "Saving",
			Number: "1234",
			Amount: 12.4,
		},
	}
	c.TplName = "banking.tpl"
}

// @router /api/transfer [post]
func (c *BankingController) Transfer() {
	var transfer models.Transfer
	json.Unmarshal(c.Ctx.Input.RequestBody, &transfer)
	fmt.Println(transfer)
	valid := validation.Validation{}
	isValid, _ := valid.Valid(&transfer)
	fmt.Println(valid.ErrorMap())

	var message string

	if isValid {
		message = "success"
	} else {
		message = "failure"
	}

	c.Ctx.WriteString(message)
}
