package search

import (
	"encoding/json"
	"os"
)

/**
用于读取json数据文件

json.NewDecoder(file).Decode(pointer of struct)
*/

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) {
	filename, _ := os.Getwd()
	filename += "/data/data.json"

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	feeds := make([]*Feed, 0)

	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err
}
