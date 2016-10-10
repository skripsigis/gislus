package controllers

import (
	"Skripsi/gisdel/webapp/commons"
	"log"
	"net/http"
	"strings"

	"Skripsi/gisdel/library/knot/knot.v1"
	"Skripsi/gisdel/library/orm"
)

type IBaseController interface {
	// not implemented anything yet
}

type BaseController struct {
	base IBaseController
	Ctx  *orm.DataContext
}

type PageInfo struct {
	PageTitle    string
	SelectedMenu string
	Breadcrumbs  map[string]string
}

func (b *BaseController) LoadBase(k *knot.WebContext) {
	k.Config.NoLog = true
	b.IsAuthenticate(k)
	b.LayoutMode(k)
}

func (b *BaseController) LayoutMode(k *knot.WebContext) {
	mode := k.Query("mode")
	if strings.ToUpper(strings.Trim(mode, "")) == "SHAREPOINT" {
		k.Config.LayoutTemplate = "_bodyonly.html"
	}
}

func (b *BaseController) IsAuthenticate(k *knot.WebContext) {
	if k.Session("userid") == nil {
		b.Redirect(k, "login", "default")
	}
	return
}

func (b *BaseController) GetUserInfo(k *knot.WebContext) interface{} {
	emp := k.Session("datamodel")
	return emp
}

func (b *BaseController) IsGranted(k *knot.WebContext, menu string) bool {
	if menu == "dashboard" {
		return true
	}

	menuaccess := k.Session("usermenus").([]string)
	for _, m := range menuaccess {
		if m == menu {
			return true
		}
	}

	return false
}

func (b *BaseController) GetMenuAccess(k *knot.WebContext) map[string]bool {
	retMenus := make(map[string]bool, 0)
	menus := commons.GetMenus()
	menuaccess := k.Session("usermenus").([]string)

	retMenus["dashboard"] = true

	found := false
	for _, v := range menus {
		for _, m := range menuaccess {
			if v == m {
				found = true
			}
		}
		retMenus[v] = found

		found = false
	}

	return retMenus
}

func (b *BaseController) LoadPartial(k *knot.WebContext, tpls []string) {
	defaultTpls := []string{"_loader.html", "_loader_alt.html"}
	if len(tpls) > 0 {
		defaultTpls = append(defaultTpls, tpls...)
	}
	k.Config.IncludeFiles = defaultTpls
}

func (b *BaseController) WriteLog(msg interface{}) {
	log.Printf("%#v\n\r", msg)
	return
}

func (b *BaseController) Redirect(k *knot.WebContext, controller string, action string) {
	log.Println("invalid session , redirecting to " + controller + "/" + action)
	http.Redirect(k.Writer, k.Request, "/"+controller+"/"+action, http.StatusTemporaryRedirect)
}

func (b *BaseController) RedirectToUrl(k *knot.WebContext, url string) {
	log.Println("invalid session , redirecting to " + url)
	http.Redirect(k.Writer, k.Request, url, http.StatusTemporaryRedirect)
}

type ResultInfo struct {
	IsError bool
	Message string
	Data    interface{}
}
