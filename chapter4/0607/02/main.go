package main

import (
	"encoding/json"
	"log"
	"os"
)

var content = `{
	"species": "はと",
	"description": "岩に止まるのが好き",
	"dimensions": {
		"height": 24,
		"width": 10
	}
}`

type Dimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Data struct {
	Species     string     `json:"species"`
	Description string     `json:"description"`
	Dimensions  Dimensions `json:"dimensions"`
}

// func main() {
// 	var data Data
// 	// 通常このようにUnmarshalを使ってパース(JSON -> 構造体)する
// 	err := json.Unmarshal([]byte(content), &data)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(&data)
// }

func main() {
	f, err := os.Open("input.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// ネットワーク通信のストリームやファイルなどを扱う時にはDecodeを使う
	var data Data
	if err := json.NewDecoder(f).Decode(&data); err != nil {
		log.Fatal(err)
	}
}
