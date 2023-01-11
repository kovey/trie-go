package filter

import (
	"fmt"
	"regexp"
)

const (
	Reg_Chinese       = "^[\u4E00-\u9FA5]+$"
	Reg_Chinese_Count = "^[\u4E00-\u9FA5]{%d,%d}$"
)

func IsChinese(text string) bool {
	ok, _ := regexp.Match(Reg_Chinese, []byte(text))
	return ok
}

func IsChineseByCount(text string, min, max int) bool {
	ok, _ := regexp.Match(fmt.Sprintf(Reg_Chinese_Count, min, max), []byte(text))
	return ok
}
