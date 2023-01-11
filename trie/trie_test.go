package trie

import "testing"

func TestTrie(t *testing.T) {
	tt := NewTrie()
	tt.Insert("你妈")
	tt.Insert("娘亲")
	tt.Insert("成人视频")
	tt.Insert("妈的")

	t.Log(tt.Replace("有人在看成人视频"))
	t.Log(tt.Replace("有人在看视频"))
	t.Log(tt.Replace("有人在看成人"))
	t.Log(tt.Replace("你妈"))
	t.Log(tt.Replace("你娘亲"))
	t.Log(tt.Replace("看你妈的"))
	t.Log(tt.Replace("妈的,看什么看"))
	t.Log(tt.Has("你妈"))
	t.Log(tt.Has("你娘亲"))
	t.Log(tt.Has("有人在看成人视频"))
	t.Log(tt.Has("有人在看成人"))
}
