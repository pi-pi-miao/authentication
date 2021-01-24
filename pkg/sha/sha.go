package sha

import (
	"authentication/pkg/base64"
	"authentication/pkg/str"
	"bytes"
	"crypto/sha256"
	"strings"
)

// sha256 data + base64 time
func GetSha256(data string,list string,permission,t string)string{
	buf := &strings.Builder{}
	buf.Grow(100)
	sum := sha256.New()
	sum.Write(str.StringToByte(data))
	buf.WriteString(base64.BaseEncode(sum.Sum(nil)))
	buf.WriteString(GetData(str.StringToByte(t),list,permission))
	d   := buf.String()
	buf.Reset()
	return d
}

func SetSha256(data string)string{
	sum := sha256.New()
	sum.Write(str.StringToByte(data))
	return base64.BaseEncode(sum.Sum(nil))
}

func GetData(t []byte,list,permission string)string{
	buf  := &bytes.Buffer{}
	buf.Grow(100)
	buf.Write(t)
	buf.WriteString(list)
	buf.WriteString(permission)
	data := buf.Bytes()
	buf.Reset()
	return base64.BaseEncode(data)
}