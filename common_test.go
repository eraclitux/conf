package cfgp

import (
	"testing"
)

// TODO use this to test flags:
//defer func() { os.Args = oldArgs }()
//os.Args = []string{"cmd", "-before", "subcmd", "-after", "args"}
// also look at  src/flag/flag_test.go

func TestParse(t *testing.T) {
	err := Parse("local.yml", nil)
	if err == nil {
		t.Fail()
	}
}
