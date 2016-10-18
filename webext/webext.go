package webext

import (
	"bufio"
	. "Skripsi/gislus/webapp/controllers"
	"fmt"
	"os"
	"strings"

	"Skripsi/gislus/library/dbox"
	_ "Skripsi/gislus/library/dbox/dbc/mongo"
	"Skripsi/gislus/library/knot/knot.v1"
	"Skripsi/gislus/library/orm"
)

var (
	wd = func() string {
		d, _ := os.Getwd()
		return d + "/"
	}()
)

func init() {
	conn, err := PrepareConnection()
	if err != nil {
		fmt.Println(err)
	}
	ctx := orm.New(conn)

	baseCont := new(BaseController)
	baseCont.Ctx = ctx

	app := knot.NewApp("gisapp")
	app.ViewsPath = wd + "views/"

	// register controllers
	app.Register(&LoginController{baseCont})
	app.Register(&DashboardController{baseCont})
	app.Register(&LiftingController{baseCont})

	app.Static("static", wd+"assets")
	app.LayoutTemplate = "_layout.html"
	knot.RegisterApp(app)
}

func PrepareConnection() (dbox.IConnection, error) {
	config := ReadConfig()
	ci := &dbox.ConnectionInfo{config["host"], config["database"], config["username"], config["password"], nil}
	c, e := dbox.NewConnection("mongo", ci)

	if e != nil {
		return nil, e
	}

	e = c.Connect()
	if e != nil {
		return nil, e
	}

	return c, nil
}

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
