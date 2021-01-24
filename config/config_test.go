package config

import "testing"

var (
	path = "/home/pi-pi/project/template/example/authentication/config/config.toml"
)
func TestInit(t *testing.T) {
	err := Init(path)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("suc")
}
