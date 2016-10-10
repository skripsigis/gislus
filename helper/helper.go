package helper

import (
	"fmt"
	"io"
	"os"

	"Skripsi/gisdel/library/dbox"
	"Skripsi/gisdel/library/knot/knot.v1"
	"Skripsi/gisdel/library/toolkit"
)

var (
	DebugMode bool
)

// type ConnectionPath struct {
// 	Path string
// }

var config_system = func() string {
	d, _ := os.Getwd()
	d += "/conf/confsystem.json"
	return d
}()

func GetPathConfig() (result map[string]interface{}) {
	result = make(map[string]interface{})

	ci := &dbox.ConnectionInfo{config_system, "", "", "", nil}
	conn, e := dbox.NewConnection("json", ci)
	if e != nil {
		return
	}

	e = conn.Connect()
	defer conn.Close()
	csr, e := conn.NewQuery().Select("*").Cursor(nil)
	if e != nil {
		return
	}
	defer csr.Close()
	data := []toolkit.M{}
	e = csr.Fetch(&data, 0, false)
	if e != nil {
		return
	}
	result["folder-path"] = data[0].GetString("folder-path")
	result["restore-path"] = data[0].GetString("restore-path")
	return
}

func CreateResult(success bool, data interface{}, message string) map[string]interface{} {
	if !success {
		fmt.Println("ERROR! ", message)
		if DebugMode {
			panic(message)
		}
	}

	return map[string]interface{}{
		"data":    data,
		"success": success,
		"message": message,
	}
}

func UploadHandler(r *knot.WebContext, filename, dstpath string, replacename string) (error, string) {
	file, handler, err := r.Request.FormFile(filename)
	if err != nil {
		return err, ""
	}
	defer file.Close()

	dstSource := dstpath + toolkit.PathSeparator + replacename
	f, err := os.OpenFile(dstSource, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err, ""
	}
	defer f.Close()
	io.Copy(f, file)

	return nil, handler.Filename
}

func UploadHandlerCopy(r *knot.WebContext, filename, dstpath string) (error, string) {
	file, handler, err := r.Request.FormFile(filename)
	if err != nil {
		return err, ""
	}
	defer file.Close()

	dstSource := dstpath + toolkit.PathSeparator + handler.Filename
	f, err := os.OpenFile(dstSource, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err, ""
	}
	defer f.Close()
	io.Copy(f, file)

	return nil, handler.Filename
}
