package lua

import (
	"golang.org/x/sys/windows"
)

func openlib() uintptr {
	handle, err := windows.LoadLibrary("luajit.dll")
	if err != nil {
		panic(err)
	}

	return uintptr(handle)
}
