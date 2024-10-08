// GoでJSONを取り扱う
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var content = `
{
	"species": "ハト",
	"description": "岩に止まるのが好き",
	"dimensions": {
		"height": 24,
		"width": 10
	}
}
`

type Dimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Data struct {
	Species     string     `json:"species"`
	Description string     `json:"description"`
	Dimensions  Dimensions `json:"dimensions"`
}

func main() {
	var data Data
	err := json.Unmarshal([]byte(content), &data) // JSON文字列をパース
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", data) // {Species:ハト Description:岩に止まるのが好き Dimensions:{Width:10 Height:24}}

	// ネットワーク通信のストリームやファイルからJSONを読み込む場合は、デコーダを使う
	f, err := os.Open("input.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var data2 Data
	err2 := json.NewDecoder(f).Decode(&data2)
	if err2 != nil {
		log.Fatal(err2)
	}

	// これによって巨大なデータであっても入力済みのデータとパース済みのstructの両方にメモリが確保されない

}
