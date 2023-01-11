package filter

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type Content struct {
	Result string `json:"result"`
}

func (c *Content) Load(path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(content, c)
}

func (c *Content) Get() []string {
	return strings.Split(c.Result[1:len(c.Result)-1], ",")
}
