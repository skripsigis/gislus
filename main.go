package main

import (
	_ "Skripsi/gisdel/webapp/webext"
	"net/http"

	"Skripsi/gisdel/library/knot/knot.v1"
)

func main() {
	app := knot.GetApp("ostroreport")
	if app == nil {
		return
	}

	routes := make(map[string]knot.FnContent, 1)
	routes["/"] = func(r *knot.WebContext) interface{} {
		http.Redirect(r.Writer, r.Request, "/login/default", http.StatusTemporaryRedirect)
		return true
	}
	knot.StartAppWithFn(app, "localhost:8017", routes)
}
