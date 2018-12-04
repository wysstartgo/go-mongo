package conf

import (
	"io/ioutil"
	"encoding/json"
)

type Conf struct {
	SourceMongoUrl string	`json:"sourceMongoUrl"`
	TargetMongoUrl string  `json:"targetMongoUrl"`
	SourceCollection string `json:"sourceCollection"`
	TargetCollection string `json:"targetCollection"`
	IsCopyDataBase  string `json:"isCopyDataBase"`
}

var GConf *Conf

func LoadConf(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	conf := Conf{}
	err = json.Unmarshal(content, &conf)
	if err != nil {
		return err
	}
	GConf = &conf
	return nil
}