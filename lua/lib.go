package lua

const (
	CoLibName   = "coroutine"
	MathLibName = "math"
	StrLibName  = "string"
	TabLibName  = "table"
	IOLibName   = "io"
	OSLibName   = "os"
	LoadLibName = "package"
	DBLibName   = "debug"
	BitLibName  = "bit"
	JitLibName  = "jit"
	FFILibName  = "ffi"
)

func OpenLibs(L State) {
	lib.openlibs(L)
}

func OpenBase(L State) int {
	return int(lib.openbase(L))
}

func OpenMath(L State) int {
	return int(lib.openmath(L))
}

func OpenString(L State) int {
	return int(lib.openstring(L))
}

func OpenTable(L State) int {
	return int(lib.opentable(L))
}

func OpenIO(L State) int {
	return int(lib.openio(L))
}

func OpenOS(L State) int {
	return int(lib.openos(L))
}

func OpenPacakge(L State) int {
	return int(lib.openpackage(L))
}

func OpenDebug(L State) int {
	return int(lib.opendebug(L))
}

func OpenBit(L State) int {
	return int(lib.openbit(L))
}

func OpenJit(L State) int {
	return int(lib.openjit(L))
}

func OpenFFI(L State) int {
	return int(lib.openffi(L))
}

func OpenStringBuffer(L State) int {
	return int(lib.openstring_buffer(L))
}

var lib struct {
	_ nocopy

	openbase          func(L State) int32 `lua:"luaopen_base"`
	openmath          func(L State) int32 `lua:"luaopen_math"`
	openstring        func(L State) int32 `lua:"luaopen_string"`
	opentable         func(L State) int32 `lua:"luaopen_table"`
	openio            func(L State) int32 `lua:"luaopen_io"`
	openos            func(L State) int32 `lua:"luaopen_os"`
	openpackage       func(L State) int32 `lua:"luaopen_package"`
	opendebug         func(L State) int32 `lua:"luaopen_debug"`
	openbit           func(L State) int32 `lua:"luaopen_bit"`
	openjit           func(L State) int32 `lua:"luaopen_jit"`
	openffi           func(L State) int32 `lua:"luaopen_ffi"`
	openstring_buffer func(L State) int32 `lua:"luaopen_string_buffer"`
	openlibs          func(L State)       `lua:"luaL_openlibs"`
}
