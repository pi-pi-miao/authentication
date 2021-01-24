package uuid

import (
	uuid "github.com/satori/go.uuid"
	"strings"
)

func GetUuid()string{
	return strings.Replace(uuid.NewV4().String(),"-","",-1)
}

func SetUuid(data string)string {
	uid,_ := uuid.FromString(data)
	return strings.Replace(uid.String(),"-","",-1)
}
