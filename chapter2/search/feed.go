package search

/**
用于读取json数据文件

json.NewDecoder(file).Decode(pointer of struct)
 */

type Feed struct {
	Name string `json:"site"`
	URI string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error){

}