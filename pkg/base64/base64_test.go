package base64

import (
	"authentication/pkg/str"
	"testing"
)

var (
	hello = "hello"
	data   string
)

func TestBaseEncode(t *testing.T) {
	data = BaseEncode(str.StringToByte(hello))
}

func TestBaseDecode(t *testing.T) {
	TestBaseEncode(t)
	src,err := BaseDecode(data)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(str.BytesToString(src))
}
