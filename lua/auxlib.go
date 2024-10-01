package lua

import (
	"unsafe"

	"github.com/ebitengine/purego"
)

// LReg is the type used to register external functions.
type LReg struct {
	Name string
	Func CFunction
}

const (
	ErrFile = ErrErr + 1 // A file cannot be open/read
)

// NewState creates a new Lua state.
func NewState() State {
	return luaL.newstate()
}

// Register opens a library.
//
// When called with libname equal to nil,
// it simply registers all functions in the list l into the table on the top of the stack.
func Register(L State, libname string, l []LReg) {
	reg := make([]lreg, len(l)+1)
	for i, fn := range l {
		name := ToCString(fn.Name)
		reg[i] = lreg{
			name: uintptr(unsafe.Pointer(&name[0])),
			fn:   purego.NewCallback(fn.Func),
		}
	}

	reg[len(l)] = lreg{
		name: 0,
		fn:   0,
	}

	luaL.register(L, libname, &reg[0])
}

// GetMetaField pushes onto the stack the field e from the metatable of the object at index obj.
// If the object does not have a metatable, or if the metatable does not have this field, returns false and pushes nothing.
func GetMetaField(L State, obj int, e string) bool {
	return luaL.getmetafield(L, int32(obj), e) == 1
}

// CallMeta calls a metamethod.
func CallMeta(L State, obj int, e string) int {
	return int(luaL.callmeta(L, int32(obj), e))
}

// TypeError generates an error with a message like the following:
//
//	location: bad argument narg to 'func' (tname expected, got rt)
//
// Where location is produced by Where, func is the name of the current function,
// and rt is the type name of the actual argument
func TypeError(L State, narg int, tname string) int {
	return int(luaL.typerror(L, int32(narg), tname))
}

// ArgError raises an error with the following message:
//
//	bad argument #<narg> to <func> (<extramsg>)
//
// Where func is retrieved from the call stack.
// This function never returns, but it is an idiom to use it in [CFunction]s as a return.
func ArgError(L State, numarg int, extramsg string) int {
	return int(luaL.argerror(L, int32(numarg), extramsg))
}

// LoadString loads a string as a Lua chunk.
//
// This function returns the same results as Load.
func LoadString(L State, s string) int {
	return int(luaL.loadstring(L, s))
}

/// Macro conversions

// ArgCheck checks whether cond is true.
// If not, raises an error with the following message,
// where func is retrieved from the call stack:
//
//	bad argument #<narg> to <func> (<extramsg>)
func ArgCheck(L State, cond bool, numarg int, extramsg string) bool {
	return cond || luaL.argerror(L, int32(numarg), extramsg) == 1
}

// CheckString checks whether the function argument numarg is a string and returns its string.
func CheckString(L State, numarg int) string {
	return luaL.checklstring(L, int32(numarg), nil)
}

func OptString(L State, numarg int, d string) string {
	return luaL.optlstring(L, int32(numarg), d, nil)
}

// CheckInt checks whether the function argument numarg is an integer and returns its number cast to an int.
func CheckInt(L State, numarg int) Integer {
	return luaL.checkinteger(L, int32(numarg))
}

func OptInt(L State, numarg int, def Integer) Integer {
	return luaL.optinteger(L, int32(numarg), def)
}

// TypeNameOf returns the name of the type of the value at the given index.
func TypeNameOf(L State, idx int) string {
	return TypeName(L, Type(L, idx))
}

// GetMetaTableFor pushes onto the stack the metatable associated with name in the registry.
func GetMetaTableFor(L State, name string) {
	GetField(L, RegistryIndex, name)
}

type lreg struct {
	name uintptr
	fn   uintptr
}

var luaL struct {
	_ nocopy

	register     func(L State, libname string, l *lreg)             `lua:"luaL_register"`
	getmetafield func(L State, obj int32, e string) int32           `lua:"luaL_getmetafield"`
	callmeta     func(L State, obj int32, e string) int32           `lua:"luaL_callmeta"`
	typerror     func(L State, narg int32, tname string) int32      `lua:"luaL_typerror"`
	argerror     func(L State, numarg int32, extramsg string) int32 `lua:"luaL_argerror"`

	checklstring func(L State, numarg int32, l *size_t) string             `lua:"luaL_checklstring"`
	optlstring   func(L State, numarg int32, def string, l *size_t) string `lua:"luaL_optlstring"`

	checknumber func(L State, numarg int32) Number             `lua:"luaL_checknumber"`
	optnumber   func(L State, numarg int32, def Number) Number `lua:"luaL_optnumber"`

	checkinteger func(L State, numarg int32) Integer              `lua:"luaL_checkinteger"`
	optinteger   func(L State, numarg int32, def Integer) Integer `lua:"luaL_optinteger"`

	checkstack func(L State, sz int32, msg string) `lua:"luaL_checkstack"`
	checktype  func(L State, narg int32, t int32)  `lua:"luaL_checktype"`
	checkany   func(L State, narg int32)           `lua:"luaL_checkany"`

	newmetatable func(L State, tname string) int32             `lua:"luaL_newmetatable"`
	checkudata   func(L State, ud int32, tname string) uintptr `lua:"luaL_checkudata"`

	where  func(L State, lvl int32)               `lua:"luaL_where"`
	error_ func(L State, fmt string, args ...any) `lua:"luaL_error"`

	newstate   func() State                  `lua:"luaL_newstate"`
	loadstring func(L State, s string) int32 `lua:"luaL_loadstring"`
}
