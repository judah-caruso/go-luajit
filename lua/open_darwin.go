//go:build darwin && !unix

package lua

import (
	"github.com/ebitengine/purego"
)

func openlib() uintptr {
	handle, err := purego.Dlopen("libluajit.dylib", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	return handle
}
