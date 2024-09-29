package lua

const (
	JitVersion   = "LuaJIT 2.1.1724232689"
	JitCopyright = "Copyright (C) 2005-2023 Mike Pall"
)

// JitMode represents a mode used to interact with the jit compiler.
type JitMode int32

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

// ProfileCallback represents a function used for profiling.
type ProfileCallback func(data uintptr, L State, samples, vmstate int32)

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

// ProfileStart starts the profiler.
func ProfileStart(L State, mode string, cb ProfileCallback, data uintptr) {
	jit.profile_start(L, mode, cb, data)
}

// ProfileStop stops the profiler.
func ProfileStop(L State) {
	jit.profile_stop(L)
}

// ProfileDumpStack allows taking stack dumps in an effecient manner.
func ProfileDumpStack(L State, fmt string, depth int, len *uint) string {
	return jit.profile_dumpstack(L, fmt, int32(depth), len)
}

var jit struct {
	setmode           func(L State, idx int32, mode JitMode) int32                 `lua:"luaJIT_setmode"`
	profile_start     func(L State, mode string, cb ProfileCallback, data uintptr) `lua:"luaJIT_profile_start"`
	profile_stop      func(L State)                                                `lua:"luaJIT_profile_stop"`
	profile_dumpstack func(L State, fmt string, depth int32, len *size_t) string   `lua:"luaJIT_profile_dumpstack"`
}
