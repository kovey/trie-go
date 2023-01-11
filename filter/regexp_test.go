package filter

import "testing"

func TestChinese(t *testing.T) {
	t.Logf("res: %t", IsChinese("你好好"))
	t.Logf("res: %t", IsChinese(""))
	t.Logf("res: %t", IsChinese("你好好 "))
	t.Logf("res: %t", IsChinese("你好好！"))
	t.Logf("res: %t", IsChinese("　你好"))
}

func TestChineseByCoun(t *testing.T) {
	t.Logf("res: %t", IsChineseByCount("你好好", 2, 4))
	t.Logf("res: %t", IsChineseByCount("你好好啊", 2, 4))
	t.Logf("res: %t", IsChineseByCount("你好", 2, 4))
	t.Logf("res: %t", IsChineseByCount("你", 2, 4))
	t.Logf("res: %t", IsChineseByCount("你真的很好", 2, 4))
	t.Logf("res: %t", IsChineseByCount("你好好 ", 2, 4))
	t.Logf("res: %t", IsChineseByCount("你好好！", 2, 4))
	t.Logf("res: %t", IsChineseByCount("　你好好", 2, 4))
}
