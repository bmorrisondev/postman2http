package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
		url := item.Request.Url.Raw
		if !strings.HasPrefix(url, "http") {
			url = "http://" + url
		}
		outString += item.Request.Method + " " + url + "\n"
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
