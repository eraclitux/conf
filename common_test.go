package cfgp

import (
	"testing"
)

func TestParse(t *testing.T) {
	err := Parse("local.yml", nil)
	if err == nil {
		t.Fail()
	}
}
