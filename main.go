package main

import (
	"archive/zip"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"./dirread"
	"./zipopen"
)

func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()
	}

	return nil
}
func ziplist(w http.ResponseWriter, r *http.Request) {
	var t zipopen.File
	var t_dir dirread.Dirtype
	data := map[string]string{}
	str := r.URL.RawQuery
	data["id"] = "0"
	data["page"] = "0"
	num := 0
	page := 0

	if strings.Index(str, "&") > 0 {
		for _, tmp := range strings.Split(str, "&") {
			if strings.Index(tmp, "=") > 0 {
				tmp2 := strings.Split(tmp, "=")
				data[tmp2[0]] = tmp2[1]

			}

		}
	}
	num, _ = strconv.Atoi(data["id"])
	page, _ = strconv.Atoi(data["page"])
	t_dir.Setup("../../../book")
	_ = t_dir.Read("/")
	if num >= len(t_dir.Data) {
		num = 0
	}
	t.ZipOpenSetup(t_dir.Data[num].RootPath + t_dir.Data[num].Name)
	t.ZipReadList()
	page--
	if page >= t.Count {
		page = 0
	}
	fmt.Fprintf(w, "%s", t.ZipRead(page))
}
func webstart() {
	fmt.Println("web server start")
	http.HandleFunc("/zip", ziplist)
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./html"))))
	http.ListenAndServe(":8080", nil)
}

func main() {
	// err := Unzip("../../../book/[大森藤ノ]ファミリアクロニクル_フレイヤ.zip", "./out")
	var t zipopen.File
	var t_dir dirread.Dirtype
	t_dir.Setup("../../../book")
	_ = t_dir.Read("/")
	t.ZipOpenSetup(t_dir.Data[0].RootPath + t_dir.Data[0].Name)
	t.ZipReadList()
	// fmt.Println(t.ZipRead(0))

	webstart()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(t_dir)
	// fmt.Println(t)
}
