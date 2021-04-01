package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// grab the file path from the args
	jsonFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	// import the json file
	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	var collection PostmanCollection

	// Unmarshal it to a type
	json.Unmarshal(byteValue, &collection)

	// fmt.Println(postmanCollection.Info.Name)
	// fmt.Println(postmanCollection.Info.PostmanId)
	// fmt.Println(postmanCollection.Info.Schema)

	// fmt.Println(string(byteValue))

	// create an out string
	outString := ""
	for idx, item := range collection.Items {
		outString += item.Request.Method + " " + item.Request.Url.Raw + "\n"
		for _, header := range item.Request.Headers {
			outString += header.Key + ": " + header.Value + "\n"
		}
		if item.Request.Body.Raw != "" {
			outString += "\n" + item.Request.Body.Raw + "\n"
		}

		// TODO: Counter to not do this on the last one
		if idx+1 != len(collection.Items) {
			outString += "\n###\n\n"
		}
	}

	// save to the disk
	f, err := os.Create(collection.Info.Name + ".http")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = f.WriteString(outString)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	// fmt.Println(l, "bytes written successfully")
	// err = f.Close()
	// if err != nil {
	// 		fmt.Println(err)
	// 		return
	// }

}

type PostmanCollection struct {
	Info  Info   `json:"info"`
	Items []Item `json:"item"`
}

type Info struct {
	PostmanId string `json:"_postman_id"`
	Name      string `json:"name"`
	Schema    string `json:"schema"`
}

type Item struct {
	Name    string  `json:"name"`
	Request Request `json:"request"`
}

type Request struct {
	Method  string   `json:"method"`
	Headers []Header `json:"header"`
	Url     Url      `json:"url"`
	Body    Body     `json:"body"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Url struct {
	Raw string `json:"raw"`
}

type Body struct {
	Raw string `json:"raw"`
}
