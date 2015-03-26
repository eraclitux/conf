// +build debug

package cfgp

import (
	"fmt"
	"os"
)

func debugPrintln(args ...interface{}) {
	// we use Stderr to not break example_test.go wich evaluate Stdout in tests
	fmt.Fprintln(os.Stderr, args...)
}
