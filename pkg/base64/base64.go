package base64

import (
	"encoding/base64"
)

func BaseEncode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func BaseDecode(src string) ([]byte,error){
	return base64.StdEncoding.DecodeString(src)
}

