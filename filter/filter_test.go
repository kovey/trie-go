package filter

import (
	"path/filepath"
	"runtime"
	"testing"
)

func TestFilter(t *testing.T) {
	if err := Load(CurrentDir()+"/keywords.txt", ","); err != nil {
		t.Fatalf("load error: %s", err)
	}

	t.Log(Replace("黑: test"))
	t.Logf("has[%t]", Has("test"))
	t.Log(Replace("你妈的"))
}

func TestFilterJson(t *testing.T) {
	if err := LoadByJson(CurrentDir() + "/keywords.json"); err != nil {
		t.Fatalf("load error: %s", err)
	}
	t.Log(Replace("我是成人视频啊啊啊"))
	t.Log(Replace("你好啊"))
	t.Logf("has[%t]", Has("成人视频"))
}

func CurrentDir() string {
	_, file, _, _ := runtime.Caller(1)
	return filepath.Dir(file)
}
