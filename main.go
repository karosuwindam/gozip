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

const ZIPPATH = "../../../book"

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
func zipdata(w http.ResponseWriter, r *http.Request) {
	var t zipopen.File
	var t_dir dirread.Dirtype
	data := map[string]string{}
	str := r.URL.RawQuery
	data["id"] = "1"
	data["page"] = "1"
	num := 1
	page := 1

	if strings.Index(str, "&") > 0 {
		for _, tmp := range strings.Split(str, "&") {
			if strings.Index(tmp, "=") > 0 {
				tmp2 := strings.Split(tmp, "=")
				data[tmp2[0]] = tmp2[1]

			}

		}
	} else if strings.Index(str, "=") > 0 {
		tmp2 := strings.Split(str, "=")
		data[tmp2[0]] = tmp2[1]

	}
	num, _ = strconv.Atoi(data["id"])
	page, _ = strconv.Atoi(data["page"])
	t_dir.Setup(ZIPPATH)
	_ = t_dir.Read("/")
	if ((num - 1) >= len(t_dir.Data)) || (num == 0) {
		num = 1
	}
	t.ZipOpenSetup(t_dir.Data[num-1].RootPath + t_dir.Data[num-1].Name)
	t.ZipReadList()
	page--
	if page >= t.Count {
		page = 0
	}
	fmt.Fprintf(w, "%s", t.ZipRead(page))
}
func ziplist(w http.ResponseWriter, r *http.Request) {
	var t zipopen.File
	var t_dir dirread.Dirtype
	str := ""
	t_dir.Setup(ZIPPATH)
	_ = t_dir.Read("/")
	for _, r := range t_dir.Data {
		str += r.RootPath + r.Name + ","
		t.ZipOpenSetup(r.RootPath + r.Name)
		t.ZipReadList()
		str += strconv.Itoa(t.Count) + "\n"
	}
	fmt.Fprintf(w, "%v", str)
}

func view(w http.ResponseWriter, r *http.Request) {
	var t zipopen.File
	var t_dir dirread.Dirtype
	var datap map[string]string
	url := r.URL.Path
	data := map[string]string{}
	datap = data
	data["id"] = "1"
	id := 0
	data["nowpage"] = "1"

	t_dir.Setup(ZIPPATH)
	_ = t_dir.Read("/")

	i := 0
	for _, str := range strings.Split(url[1:], "/") {
		if (i == 1) && (str != "") {
			tmp, _ := strconv.Atoi(str)
			if tmp > 0 {
				if len(t_dir.Data) >= tmp {
					data["id"] = str
					id = tmp - 1
				}
			}
		}
		if (i == 2) && (str != "") {
			tmp, _ := strconv.Atoi(str)
			if tmp > 0 {
				data["nowpage"] = str
			}
		}
		println(str)
		i++
	}
	t.ZipOpenSetup(t_dir.Data[id].RootPath + t_dir.Data[id].Name)
	t.ZipReadList()
	data["title"] = t_dir.Data[id].Name[1:]
	data["pagemax"] = strconv.Itoa(t.Count)
	// output := ConvertData(ReadHtml("html/comic/view.html"), datap)
	output := ConvertData(ReadHtml("html/view.html"), datap)
	fmt.Fprintf(w, output)
	// fmt.Fprintf(w, "id=%vnowpage=%vpagemax=%v", data["id"], data["nowpage"], data["pagemax"])
}

func webstart() {
	fmt.Println("web server start")
	http.HandleFunc("/zip", zipdata)
	http.HandleFunc("/ziplist", ziplist)
	http.HandleFunc("/view/", view)
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
