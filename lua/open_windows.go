//go:build windows

package lua

import (
	"golang.org/x/sys/windows"
)

func openlib() uintptr {
	handle, err := windows.LoadLibrary("libluajit.dll")
	if err != nil {
		panic(err)
	}

	return uintptr(handle)
}
