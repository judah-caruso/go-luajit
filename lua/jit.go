package lua

const (
	JitVersion   = "LuaJIT 2.1.1724232689"
	JitCopyright = "Copyright (C) 2005-2023 Mike Pall"
)

type (
	JitMode int32
)

const (
	ModeEngine JitMode = iota
	ModeDebug
	ModeFunc
	ModeAllFunc
	ModeAllSubFunc
	ModeTrace
	ModeWrapCFunc = JitMode(0x10)
	ModeMax       = ModeWrapCFunc + 1
)

const (
	ModeOff   = JitMode(0x0000)
	ModeOn    = JitMode(0x0100)
	ModeFlush = JitMode(0x0200)
	ModeMask  = JitMode(0x00ff)
)

// SetMode allows control of the VM.
//
// 'idx' is expected to be 0 or a stack index.
//
// 'mode' specifies the mode, which is 'or'ed with a flag. The flag can be:
//   - ModeOff to turn a feature off
//   - ModeOn to turn a feature on
//   - ModeFlush to flush cached code.
func SetMode(L State, idx int, mode JitMode) bool {
	return jit.setmode(L, int32(idx), mode) == 1
}

var jit struct {
	setmode func(L State, idx int32, mode JitMode) int32 `lua:"luaJIT_setmode"`
}
