package main

import (
	"fmt"

	"github.com/judah-caruso/go-luajit/lua"
)

func main() {
	state := lua.NewState()
	defer lua.Close(state)

	lua.OpenLibs(state)

	lua.LoadString(state, `return unpack(...)[1]()`)

	lua.PushClosure(state, func(L lua.State) int {
		fmt.Println("hello from go!")
		lua.PushString(L, "it works!")
		return 1
	}, 0)

	lua.Call(state, 0, 1)

	fmt.Println("go function return type:", lua.TypeNameOf(state, -1))
	fmt.Println("go function returned:", lua.ToString(state, -1))
}
