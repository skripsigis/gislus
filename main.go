package main

import (
	_ "Skripsi/gislus/webapp/webext"
	"net/http"
		"bufio"
	"fmt"
	"os"
	"strings"

	"Skripsi/gislus/library/knot/knot.v1"
)

func main() {
	app := knot.GetApp("gisapp")
	if app == nil {
		return
	}
	config := ReadConfig()

	routes := make(map[string]knot.FnContent, 1)
	routes["/"] = func(r *knot.WebContext) interface{} {
		http.Redirect(r.Writer, r.Request, "/login/default", http.StatusTemporaryRedirect)
		return true
	}
	knot.StartAppWithFn(app, config["Url"], routes)
}
var (
	wd = func() string {
		d, _ := os.Getwd()
		return d + "/"
	}()
)

func ReadConfig() map[string]string {
	ret := make(map[string]string)
	file, err := os.Open(wd + "conf/app.conf")
	if err == nil {
		defer file.Close()

		reader := bufio.NewReader(file)
		for {
			line, _, e := reader.ReadLine()
			if e != nil {
				break
			}

			sval := strings.Split(string(line), "=")
			ret[sval[0]] = sval[1]
		}
	} else {
		fmt.Println(err.Error())
	}

	return ret
}