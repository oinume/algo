// +build tools

package tools

import (
	_ "github.com/client9/misspell"
	_ "github.com/kisielk/errcheck"
	_ "github.com/kouzoh/wrench"
	_ "golang.org/x/lint/golint"
	_ "golang.org/x/tools/cmd/goimports"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
