package conf

import "testing"

func TestConf_LoadFromPath(t *testing.T) {
	c := New()
	t.Log(c.LoadFromPath("../example/aiko.yml.example"))
}