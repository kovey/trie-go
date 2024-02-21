package filter

import (
	"encoding/json"
	"os"
	"strings"
)

type Content struct {
	Result string `json:"result"`
}

func (c *Content) Load(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(content, c)
}

func (c *Content) Get() []string {
	return strings.Split(c.Result[1:len(c.Result)-1], ",")
}
