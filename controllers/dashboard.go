package controllers

import (
	"Skripsi/gislus/library/knot/knot.v1"
)

type DashboardController struct {
	*BaseController
}

func (c *DashboardController) Default(k *knot.WebContext) interface{} {
	c.LoadBase(k)

	k.Config.OutputType = knot.OutputTemplate

	infos := PageInfo{}
	infos.PageTitle = "Dashboard"
	infos.SelectedMenu = "dashboard"
	infos.Breadcrumbs = make(map[string]string, 0)

	return infos
}
