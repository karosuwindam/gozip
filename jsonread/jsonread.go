package jsonread

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const FILEPATH = "jsonread/conf.json"

type KaijiJson struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

type Kaiji struct {
	Data []KaijiJson
}

func (t *Kaiji) Readdata() {
	byte, err := ioutil.ReadFile(FILEPATH)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(byte, &t.Data); err != nil {
		log.Fatal(err)
	}
	return
}

func (t *Kaiji) AddData(data *KaijiJson) {
	tmp := new(KaijiJson)
	if data.Id == 0 {
		tmp.Id = len(t.Data) + 1
	}
	tmp.Title = data.Title
	tmp.Text = data.Text
	t.Data = append(t.Data, *tmp)
	return
}

func (t Kaiji) JsonEcod() string {
	tmp, err := json.Marshal(t.Data)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	output := string(tmp)
	return output
}
func (t Kaiji) WriteData() {
	tmp := t.JsonEcod()
	fp, err := os.Create(FILEPATH)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer fp.Close()
	fmt.Fprintf(fp, "%s", tmp)
	return
}

/*
func main() {
	test := new(Kaiji)
	test.Readdata()
	tmptype := new(KaijiJson)
	tmptype.Text = "aaaaaa"
	tmptype.Title = "bbbb"
	test.AddData(tmptype)
	fmt.Println(test.Data)
	fmt.Println(test.JsonEcod())
	test.WriteData()
}
*/
