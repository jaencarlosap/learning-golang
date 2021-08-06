package db

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func DB(response interface{}) {
	jsonFile, _ := os.Open("db/db.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &response)
}
