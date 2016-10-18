package controllers

import (
	"Skripsi/gislus/library/knot/knot.v1"
)

type LiftingController struct {
	*BaseController
}

func (c *LiftingController) Default(k *knot.WebContext) interface{} {
	c.LoadBase(k)

	k.Config.OutputType = knot.OutputTemplate

	infos := PageInfo{}
	infos.PageTitle = "Lifting Analytic"
	infos.SelectedMenu = "lifting"
	infos.Breadcrumbs = make(map[string]string, 0)

	return infos
}
