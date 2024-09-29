package lua

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/ebitengine/purego"
)

const (
	Version    = "Lua 5.1"
	Release    = "Lua 5.1.4"
	VersionNum = 501
	Copyright  = "Copyright (C) 1994-2008 Lua.org, PUC-Rio"
	Authors    = "R. Ierusalimschy, L. H. de Figueiredo & W. Celes"
)

// Mark for precompiled code
const Signature = "\033Lua"

type (
	State     uintptr
	Integer   = int64
	Number    = float64
	CFunction func(L State) (nresults int32)
)

const (
	MultRet       = -1
	RegistryIndex = -10000
	EnvironIndex  = -10001
	GlobalsIndex  = -10002
)

const (
	StatusOk    = 0
	StatusYield = 1
	ErrRun      = 2
	ErrSyntax   = 3
	ErrMem      = 4
	ErrErr      = 5
)

type T int32

const (
	TNone          = T(-1)
	TNil           = T(0)
	TBoolean       = T(1)
	TLightUserData = T(2)
	TNumber        = T(3)
	TString        = T(4)
	TTable         = T(5)
	TFunction      = T(6)
	TUserData      = T(7)
	TThread        = T(8)
)

const (
	MinStack = 20
)

const (
	GCStop       = 0
	GCRestart    = 1
	GCCollect    = 2
	GCCount      = 3
	GCCountB     = 4
	GCStep       = 5
	GCSetPause   = 6
	GCSetStepMul = 7
	GCIsRunning  = 9
)

// Open creates a new Lua state.
func Open() State {
	return luaL.newstate()
}

// Close destroys all objects in the given Lua state (calling the corresponding garbage-collection metamethods, if any)
// and frees all dynamic memory used by this state
func Close(L State) {
	lua.close(L)
}

// NewThread creates a new thread, pushes it on the stack, and returns a new Lua state that represents this new thread.
// The new state returned by this function shares with the original state all global objects (such as tables), but has an independent execution stack.
//
// There is no explicit function to close or to destroy a thread. Threads are subject to garbage collection, like any Lua object.
func NewThread(L State) State {
	return lua.newthread(L)
}

// GetTop returns the index of the top element in the stack.
//
// Because indices start at 1, this result is equal to the number of elements in the stack (and so 0 means an empty stack).
func GetTop(L State) int {
	return int(lua.gettop(L))
}

// SetTop accepts any acceptable index, or 0, and sets the stack top to this index.
// If the new top is larger than the old one, then the new elements are filled with nil.
// If index is 0, then all stack elements are removed.
func SetTop(L State, idx int) {
	lua.settop(L, int32(idx))
}

// PushValue pushes a copy of the element at the given valid index onto the stack.
func PushValue(L State, idx int) {
	lua.pushvalue(L, int32(idx))
}

// Remove removes the element at the given valid index,
// shifting down the elements above this index to fill the gap.
// Cannot be called with a pseudo-index, because a pseudo-index is not an actual stack position.
func Remove(L State, idx int) {
	lua.remove(L, int32(idx))
}

// Insert moves the top element into the given valid index,
// shifting up the elements above this index to open space.
// Cannot be called with a pseudo-index, because a pseudo-index is not an actual stack position.
func Insert(L State, idx int) {
	lua.insert(L, int32(idx))
}

// Replace moves the top element into the given position (and pops it),
// without shifting any element (therefore replacing the value at the given position).
func Replace(L State, idx int) {
	lua.replace(L, int32(idx))
}

// CheckStack ensures that there are at least extra free stack slots in the stack.
// It returns false if it cannot grow the stack to that size.
// This function never shrinks the stack; if the stack is already larger than the new size,
// it is left unchanged.
func CheckStack(L State, sz int) int {
	return int(lua.checkstack(L, int32(sz)))
}

// XMove exchanges values between different threads of the same global state.
//
// This function pops n values from the stack from, and pushes them onto the stack to.
func XMove(from, to State, n int) {
	lua.xmove(from, to, int32(n))
}

// IsNumber returns true if the value at the given acceptable index is a number or a string convertible to a number, and false otherwise.
func IsNumber(L State, idx int) bool {
	return lua.isnumber(L, int32(idx)) == 1
}

// IsString returns true if the value at the given acceptable index is a string or a number (which is always convertible to a string), and false otherwise.
func IsString(L State, idx int) bool {
	return lua.isstring(L, int32(idx)) == 1
}

// IsCFunction returns true if the value at the given acceptable index is a C function, and false otherwise.
func IsCFunction(L State, idx int) bool {
	return lua.iscfunction(L, int32(idx)) == 1
}

// IsUserData returns true if the value at the given acceptable index is a userdata (either full or light), and false otherwise.
func IsUserData(L State, idx int) bool {
	return lua.isuserdata(L, int32(idx)) == 1
}

// Type returns the type of the value in the given acceptable index, or TNone for a non-valid index (that is, an index to an "empty" stack position).
func Type(L State, idx int) T {
	return T(lua.type_(L, int32(idx)))
}

// TypeName returns the name of the type encoded by the value tp.
func TypeName(L State, tp T) string {
	return lua.typename(L, tp)
}

// Equals returns true if the two values in acceptable indices index1 and index2 are equal,
// following the semantics of the Lua == operator (that is, may call metamethods).
// Otherwise returns false.
//
// Also returns false if any of the indices is non valid.
func Equal(L State, index1, index2 int) bool {
	return lua.equal(L, int32(index1), int32(index2)) == 1
}

// RawEqual returns true if the two values in acceptable indices index1 and index2 are primitively equal (that is, without calling metamethods).
// Otherwise returns false.
//
// Also returns false if any of the indices are non valid.
func RawEqual(L State, index1, index2 int) bool {
	return lua.rawequal(L, int32(index1), int32(index2)) == 1
}

// LessThan returns true if the value at acceptable index index1 is smaller than the value at acceptable index index2,
// following the semantics of the Lua < operator (that is, may call metamethods).
// Otherwise returns false.
//
// Also returns false if any of the indices is non valid.
func LessThan(L State, index1, index2 int) bool {
	return lua.lessthan(L, int32(index1), int32(index2)) == 1
}

// ToNumber converts the Lua value at the given acceptable index to the type Number.
func ToNumber(L State, idx int) Number {
	return lua.tonumber(L, int32(idx))
}

// ToInteger converts the Lua value at the given acceptable index to the type Integer.
func ToInteger(L State, idx int) Integer {
	return lua.tointeger(L, int32(idx))
}

// ToBoolean converts the Lua value at the given acceptable index to a boolean value.
func ToBoolean(L State, idx int) bool {
	return lua.toboolean(L, int32(idx))
}

// ToString converts the Lua value at the given acceptable index to a string value.
func ToString(L State, idx int) string {
	return ToLString(L, idx, nil)
}

func ToLString(L State, idx int, len *uint) string {
	return lua.tolstring(L, int32(idx), len)
}

func ObjLen(L State, idx int) int {
	return int(lua.objlen(L, int32(idx)))
}

func ToUserData(L State, idx int) uintptr {
	return lua.touserdata(L, int32(idx))
}

func ToThread(L State, idx int) State {
	return lua.tothread(L, int32(idx))
}

func ToPointer(L State, idx int) uintptr {
	return lua.topointer(L, int32(idx))
}

func ToCFunction(L State, idx int) CFunction {
	return lua.tocfunction(L, int32(idx))
}

func PushNil(L State) {
	lua.pushnil(L)
}

func PushNumber(L State, n Number) {
	lua.pushnumber(L, n)
}

func PushInteger(L State, n Integer) {
	lua.pushinteger(L, n)
}

func PushLString(L State, s string, l int) {
	lua.pushlstring(L, s, size_t(l))
}

func PushString(L State, s string) {
	lua.pushstring(L, s)
}

func PushBoolean(L State, b bool) {
	var v int32 = 0
	if b {
		v = 1
	}
	lua.pushboolean(L, v)
}

func PushLightUserData(L State, p uintptr) {
	lua.pushlightuserdata(L, p)
}

func PushThread(L State) int {
	return int(lua.pushthread(L))
}

func PushClosure(L State, fn CFunction, n int) {
	lua.pushcclosure(L, fn, int32(n))
}

func GetTable(L State, idx int) {
	lua.gettable(L, int32(idx))
}

func GetField(L State, idx int, k string) {
	lua.getfield(L, int32(idx), k)
}

func RawGet(L State, idx int) {
	lua.rawget(L, int32(idx))
}

func RawGetI(L State, idx, n int) {
	lua.rawgeti(L, int32(idx), int32(n))
}

func CreateTable(L State, narr int, nrec int) {
	lua.createtable(L, int32(narr), int32(nrec))
}

func NewUserData(L State, sz int) uintptr {
	return lua.newuserdata(L, size_t(sz))
}

func GetMetaTable(L State, objindex int) int {
	return int(lua.getmetatable(L, int32(objindex)))
}

func GetFEnv(L State, idx int) {
	lua.getfenv(L, int32(idx))
}

func SetTable(L State, idx int) {
	lua.settable(L, int32(idx))
}

func SetField(L State, idx int, k string) {
	lua.setfield(L, int32(idx), k)
}

func RawSet(L State, idx int) {
	lua.rawset(L, int32(idx))
}

func RawSetI(L State, idx, n int) {
	lua.rawseti(L, int32(idx), int32(n))
}

func SetMetaTable(L State, objindex int) int {
	return int(lua.setmetatable(L, int32(objindex)))
}

func SetFEnv(L State, idx int) int {
	return int(lua.setfenv(L, int32(idx)))
}

func Call(L State, nargs, nresults int) {
	lua.call(L, int32(nargs), int32(nresults))
}

func PCall(L State, nargs, nresults, errfunc int) int {
	return int(lua.pcall(L, int32(nargs), int32(nresults), int32(errfunc)))
}

func Yield(L State, nresults int) int {
	return int(lua.yield(L, int32(nresults)))
}

func Resume(L State, narg int) int {
	return int(lua.resume(L, int32(narg)))
}

func Status(L State) int {
	return int(lua.status(L))
}

func GC(L State, what, data int) int {
	return int(lua.gc(L, int32(what), int32(data)))
}

func Error(L State) int {
	return int(lua.error_(L))
}

func Next(L State, idx int) int {
	return int(lua.next(L, int32(idx)))
}

func Concat(L State, n int) {
	lua.concat(L, int32(n))
}

// Macro conversions
// -----------------

func UpvalueIndex(i int) int {
	return int(int32(GlobalsIndex) - int32(i))
}

func Pop(L State, n int) {
	lua.settop(L, -int32(n)-1)
}

func NewTable(L State) {
	lua.createtable(L, 0, 0)
}

func Strlen(L State, i int) int {
	return int(lua.objlen(L, int32(i)))
}

func IsFunction(L State, n int) bool {
	return lua.type_(L, int32(n)) == TFunction
}

func IsTable(L State, n int) bool {
	return lua.type_(L, int32(n)) == TTable
}

func IsLightUserData(L State, n int) bool {
	return lua.type_(L, int32(n)) == TLightUserData
}

func IsNil(L State, n int) bool {
	return lua.type_(L, int32(n)) == TNil
}

func IsBoolean(L State, n int) bool {
	return lua.type_(L, int32(n)) == TBoolean
}

func IsThread(L State, n int) bool {
	return lua.type_(L, int32(n)) == TThread
}

func IsNone(L State, n int) bool {
	return lua.type_(L, int32(n)) == TNone
}

func IsNoneOrNil(L State, n int) bool {
	return lua.type_(L, int32(n)) <= 0
}

func SetGlobal(L State, s string) {
	lua.setfield(L, GlobalsIndex, s)
}

func GetGlobal(L State, s string) {
	lua.getfield(L, GlobalsIndex, s)
}

func GetRegistry(L State) {
	lua.pushvalue(L, RegistryIndex)
}

func GetGCCount(L State) int {
	return int(lua.gc(L, GCCount, 0))
}

// should this be int, uint, or uintptr?
type size_t = uint

var lua struct {
	_ nocopy

	close     func(L State)       `lua:"lua_close"`
	newthread func(L State) State `lua:"lua_newthread"`

	gettop    func(L State) int32      `lua:"lua_gettop"`
	settop    func(L State, idx int32) `lua:"lua_settop"`
	pushvalue func(L State, idx int32) `lua:"lua_pushvalue"`

	remove     func(L State, idx int32)      `lua:"lua_remove"`
	insert     func(L State, idx int32)      `lua:"lua_insert"`
	replace    func(L State, idx int32)      `lua:"lua_replace"`
	checkstack func(L State, sz int32) int32 `lua:"lua_checkstack"`
	xmove      func(from, to State, n int32) `lua:"lua_xmove"`

	isnumber    func(L State, idx int32) int32 `lua:"lua_isnumber"`
	isstring    func(L State, idx int32) int32 `lua:"lua_isstring"`
	iscfunction func(L State, idx int32) int32 `lua:"lua_iscfunction"`
	isuserdata  func(L State, idx int32) int32 `lua:"lua_isuserdata"`
	type_       func(L State, idx int32) T     `lua:"lua_type"`
	typename    func(L State, tp T) string     `lua:"lua_typename"`

	equal    func(L State, idx1 int32, idx2 int32) int32 `lua:"lua_equal"`
	rawequal func(L State, idx1 int32, idx2 int32) int32 `lua:"lua_rawequal"`
	lessthan func(L State, idx1 int32, idx2 int32) int32 `lua:"lua_lessthan"`

	tonumber    func(L State, idx int32) Number              `lua:"lua_tonumber"`
	tointeger   func(L State, idx int32) Integer             `lua:"lua_tointeger"`
	toboolean   func(L State, idx int32) bool                `lua:"lua_toboolean"`
	tolstring   func(L State, idx int32, len *size_t) string `lua:"lua_tolstring"`
	objlen      func(L State, idx int32) size_t              `lua:"lua_objlen"`
	touserdata  func(L State, idx int32) uintptr             `lua:"lua_touserdata"`
	tothread    func(L State, idx int32) State               `lua:"lua_tothread"`
	topointer   func(L State, idx int32) uintptr             `lua:"lua_topointer"`
	tocfunction func(L State, idx int32) CFunction           `lua:"lua_tocfunction"`

	pushnil           func(L State)                        `lua:"lua_pushnil"`
	pushnumber        func(L State, n Number)              `lua:"lua_pushnumber"`
	pushinteger       func(L State, n Integer)             `lua:"lua_pushinteger"`
	pushlstring       func(L State, s string, l size_t)    `lua:"lua_pushlstring"`
	pushstring        func(L State, s string)              `lua:"lua_pushstring"`
	pushboolean       func(L State, b int32)               `lua:"lua_pushboolean"`
	pushlightuserdata func(L State, p uintptr)             `lua:"lua_pushlightuserdata"`
	pushthread        func(L State) int32                  `lua:"lua_pushthread"`
	pushcclosure      func(L State, fn CFunction, n int32) `lua:"lua_pushcclosure"`
	// LUA_API const char *(lua_pushvfstring) (lua_State *L, const char *fmt, va_list argp);
	// LUA_API const char *(lua_pushfstring) (lua_State *L, const char *fmt, ...);

	gettable     func(L State, idx int32)              `lua:"lua_gettable"`
	getfield     func(L State, idx int32, k string)    `lua:"lua_getfield"`
	rawget       func(L State, idx int32)              `lua:"lua_rawget"`
	rawgeti      func(L State, idx int32, n int32)     `lua:"lua_rawgeti"`
	createtable  func(L State, narr int32, nrec int32) `lua:"lua_createtable"`
	newuserdata  func(L State, sz size_t) uintptr      `lua:"lua_newuserdata"`
	getmetatable func(L State, objindex int32) int32   `lua:"lua_getmetatable"`
	getfenv      func(L State, idx int32)              `lua:"lua_getfenv"`

	settable     func(L State, idx int32)            `lua:"lua_settable"`
	setfield     func(L State, idx int32, k string)  `lua:"lua_setfield"`
	rawset       func(L State, idx int32)            `lua:"lua_rawset"`
	rawseti      func(L State, idx int32, n int32)   `lua:"lua_rawseti"`
	setmetatable func(L State, objindex int32) int32 `lua:"lua_setmetatable"`
	setfenv      func(L State, idx int32) int32      `lua:"lua_setfenv"`

	call  func(L State, nargs int32, nresults int32)                      `lua:"lua_call"`
	pcall func(L State, nargs int32, nresults int32, errfunc int32) int32 `lua:"lua_pcall"`

	yield  func(L State, nresults int32) int32 `lua:"lua_yield"`
	resume func(L State, narg int32) int32     `lua:"lua_resume"`
	status func(L State) int32                 `lua:"lua_status"`

	gc func(L State, what int32, data int32) int32 `lua:"lua_gc"`

	error_ func(L State) int32            `lua:"lua_error"`
	next   func(L State, idx int32) int32 `lua:"lua_next"`
	concat func(L State, n int32)         `lua:"lua_concat"`
}

// @todo(judah): maybe it's just better to let the user pass a library handle in

func init() {
	handle := openlib()
	bindFuncPointers(&lua, handle)
	bindFuncPointers(&luaL, handle)
	bindFuncPointers(&lib, handle)
	bindFuncPointers(&jit, handle)
}

func bindFuncPointers(structPtr any, handle uintptr) {
	t := reflect.TypeOf(structPtr).Elem()
	v := reflect.ValueOf(structPtr)
	for _, field := range reflect.VisibleFields(t) {
		if field.Name == "_" {
			continue
		}

		tag := field.Tag.Get("lua")
		if len(tag) == 0 {
			panic(fmt.Sprintf("field %q did not have a 'lua' tag", field.Name))
		}

		fnfield := v.Elem().FieldByIndex(field.Index)
		fn := reflect.NewAt(fnfield.Type(), unsafe.Pointer(fnfield.UnsafeAddr())).Elem()
		purego.RegisterLibFunc(fn.Addr().Interface(), handle, tag)
	}
}

type nocopy struct{}

func (*nocopy) Lock()   {}
func (*nocopy) UnLock() {}
