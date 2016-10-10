package controllers

import (
	"Skripsi/gisdel/library/knot/knot.v1"
)

type UserController struct {
	*BaseController
}


func (m *UserController) User(k *knot.WebContext) interface{} {
	m.LoadBase(k)
	m.LoadPartial(k, []string{})

	k.Config.OutputType = knot.OutputTemplate

	infos := PageInfo{}
	infos.PageTitle = "User Management"
	infos.SelectedMenu = "user"
	infos.Breadcrumbs = make(map[string]string, 0)

	return infos
}
