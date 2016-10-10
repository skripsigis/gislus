package controllers

import (
	"Skripsi/gisdel/webapp/commons"
	. "Skripsi/gisdel/webapp/models"

	"Skripsi/gisdel/library/dbox"
	"Skripsi/gisdel/library/knot/knot.v1"
	tk "Skripsi/gisdel/library/toolkit"
)

type LoginController struct {
	*BaseController
}

func (c *LoginController) Default(k *knot.WebContext) interface{} {
	c.LoadPartial(k, []string{})
	k.Config.NoLog = true

	k.Config.OutputType = knot.OutputTemplate
	k.Config.LayoutTemplate = "_bodyonly.html"
	return ""
}

func (c *LoginController) DoLogin(k *knot.WebContext) interface{} {

	k.Config.OutputType = knot.OutputJson
	k.Config.NoLog = true

	isLogged := false
	DefaultPage := ""
	msg := ""
	p := struct {
		UserId   string
		Password string
	}{}
	e := k.GetPayload(&p)
	if e != nil {
		c.WriteLog(e)
		msg = "Error: " + e.Error()
	}

	query := tk.M{}.Set("where", dbox.Eq("userid", p.UserId))
	csr, err := c.Ctx.Find(new(UserModel), query)
	defer csr.Close()
	if err != nil {
		return err.Error()
	}
	results := make([]UserModel, 0)
	err = csr.Fetch(&results, 0, false)
	if err != nil {
		return err.Error()
	}
	if len(results) > 0 {
		resEmp := results[0]
		if commons.GetMD5Hash(p.Password) == resEmp.Password {

			k.SetSession("userid", resEmp.UserId)
			k.SetSession("username", resEmp.UserName)
			k.SetSession("datamodel", resEmp)
			k.SetSession("referer", "OSTROPMS")

			isLogged = true
		} else {
			return "Invalid Employee ID or password!"
		}
	} else {
		return "Invalid Employee ID or password!"
	}

	return tk.M{}.Set("IsLogged", isLogged).Set("Message", msg).Set("DefaultPage", DefaultPage)
}

func (c *LoginController) DoLogout(k *knot.WebContext) interface{} {
	k.Config.NoLog = true
	k.Config.OutputType = knot.OutputNone

	referer := k.Session("referer")

	k.SetSession("userid", nil)
	k.SetSession("username", nil)
	k.SetSession("datamodel", nil)
	k.SetSession("userrole", nil)
	k.SetSession("usermenus", nil)
	k.SetSession("referer", nil)

	if referer == "SHAREPOINT" {
		c.RedirectToUrl(k, "http://portal.ostro.in/sites/Portal/")
	} else {
		c.Redirect(k, "login", "default")
	}

	return true
}
