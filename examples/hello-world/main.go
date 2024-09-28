package main

import (
	"github.com/judah-caruso/go-luajit/lua"
)

func main() {
	L := lua.NewState()
	defer lua.Close(L)

	lua.OpenLibs(L)
	lua.LoadString(L, `print("hello from luajit!")`)
	lua.Call(L, 0, 0)
}
