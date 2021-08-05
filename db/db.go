package db

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func DB() interface{} {
	os.Open("users.json")
	// Open our jsonFile
	jsonFile, _ := os.Open("users.json")

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	return result
}
