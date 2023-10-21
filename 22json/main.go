package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int    `json:"paise"`
	Platform string
	Password string   `json:"-"`
	tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("json")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"React js", 299, "learncodeonline.in", "abc123", []string{"web-dev", "js"}},
		{"svelte js", 199, "learncodeonline.in", "bca123", []string{"web-dev", "js", "svelte"}},
		{"vite js", 399, "learncodeonline.in", "fdsa23", []string{"web-dev", "js", "fast"}},
		{"hue hue hue js", 399, "learncodeonline.in", "fdsa23", nil},
	}

	// finalJson, err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "learncode.in",
		"tags": ["web-dev","js"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("json was not valid")
	}

	// some cases wher you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v", k, v)
	}

}
