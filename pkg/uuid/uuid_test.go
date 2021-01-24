package uuid

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestGetUuid(t *testing.T) {
	u := uuid.NewV4()
	fmt.Println(u)
}
