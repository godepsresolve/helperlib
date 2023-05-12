package helperlib

import (
	"fmt"
	"github.com/godepsresolve/corelib"
)

func ProvideAssistance(input string) string {
	return fmt.Sprintf("%s@v%s -> %s", pkg, version, corelib.Format(input))
}
