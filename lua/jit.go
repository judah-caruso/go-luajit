package lua

const (
	JitVersion   = "LuaJIT 2.1.1724232689"
	JitCopyright = "Copyright (C) 2005-2023 Mike Pall"
)

const (
	ModeEngine     = 0
	ModeDebug      = 1
	ModeFunc       = 2
	ModeAllFunc    = 3
	ModeAllSubFunc = 4
	ModeTrace      = 5
	ModeWrapCFunc  = 6
	ModeMax        = 7
)

const (
	ModeOff   = 0x0000
	ModeOn    = 0x0100
	ModeFlush = 0x0200
	ModeMask  = 0x00ff
)

var jit struct {
	setmode func(L State, idx, mode int32) int32 `lua:"luaJIT_setmode"`
}
