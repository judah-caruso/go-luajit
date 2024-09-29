package lua

const (
	CoroutineLibName = "coroutine"
	MathLibName      = "math"
	StringLibName    = "string"
	TableLibName     = "table"
	IoLibName        = "io"
	OsLibName        = "os"
	LoadLibName      = "package"
	DbLibName        = "debug"
	BitLibName       = "bit"
	JitLibName       = "jit"
	FfiLibName       = "ffi"
)

// OpenLibs opens all standard Lua libraries into the given state.
func OpenLibs(L State) {
	lib.open_libs(L)
}

// OpenBase opens the base library into the given state.
func OpenBase(L State) int {
	return int(lib.open_base(L))
}

// OpenMath opens the math library into the given state.
func OpenMath(L State) int {
	return int(lib.open_math(L))
}

// OpenString opens the string library into the given state.
func OpenString(L State) int {
	return int(lib.open_string(L))
}

// OpenTable opens the table library into the given state.
func OpenTable(L State) int {
	return int(lib.open_table(L))
}

// OpenIo opens the io library into the given state.
func OpenIo(L State) int {
	return int(lib.open_io(L))
}

// OpenOs opens the os library into the given state.
func OpenOs(L State) int {
	return int(lib.open_os(L))
}

// OpenPackage opens the package library into the given state.
func OpenPacakge(L State) int {
	return int(lib.open_package(L))
}

// OpenDebug opens the debug library into the given state.
func OpenDebug(L State) int {
	return int(lib.open_debug(L))
}

// OpenBit opens the bit library into the given state.
func OpenBit(L State) int {
	return int(lib.open_bit(L))
}

// OpenJit opens the jit library into the given state.
func OpenJit(L State) int {
	return int(lib.open_jit(L))
}

// OpenFfi opens the ffi library into the given state.
func OpenFfi(L State) int {
	return int(lib.open_ffi(L))
}

// OpenStringBuilder opens the string builder library into the given state.
func OpenStringBuffer(L State) int {
	return int(lib.open_string_buffer(L))
}

var lib struct {
	_ nocopy

	open_base          func(L State) int32 `lua:"luaopen_base"`
	open_math          func(L State) int32 `lua:"luaopen_math"`
	open_string        func(L State) int32 `lua:"luaopen_string"`
	open_table         func(L State) int32 `lua:"luaopen_table"`
	open_io            func(L State) int32 `lua:"luaopen_io"`
	open_os            func(L State) int32 `lua:"luaopen_os"`
	open_package       func(L State) int32 `lua:"luaopen_package"`
	open_debug         func(L State) int32 `lua:"luaopen_debug"`
	open_bit           func(L State) int32 `lua:"luaopen_bit"`
	open_jit           func(L State) int32 `lua:"luaopen_jit"`
	open_ffi           func(L State) int32 `lua:"luaopen_ffi"`
	open_string_buffer func(L State) int32 `lua:"luaopen_string_buffer"`
	open_libs          func(L State)       `lua:"luaL_openlibs"`
}
