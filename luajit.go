package luajit

import (
	"github.com/judah-caruso/go-luajit/lua"
)

type State lua.State

func NewState() State {
	return State(lua.NewState())
}

func (s State) Close() {
	lua.Close(lua.State(s))
}
