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
	Signature  = "\033Lua"
)

type (
	State     uintptr
	Integer   = int64
	Number    = float64
	CFunction = func(L State) (nresults int)
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

const (
	TNone          = -1
	TNil           = 0
	TBoolean       = 2
	TLightUserData = 2
	TNumber        = 3
	TString        = 4
	TTable         = 5
	TFunction      = 6
	TUserData      = 7
	TThread        = 8
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

func Open() State {
	return luaL.newstate()
}

func Close(L State) {
	lua.close(L)
}

func NewThread(L State) State {
	return lua.newthread(L)
}

func GetTop(L State) int {
	return int(lua.gettop(L))
}

func SetTop(L State, idx int) {
	lua.settop(L, int32(idx))
}

func PushValue(L State, idx int) {
	lua.pushvalue(L, int32(idx))
}

func Remove(L State, idx int) {
	lua.remove(L, int32(idx))
}

func Insert(L State, idx int) {
	lua.insert(L, int32(idx))
}

func Replace(L State, idx int) {
	lua.replace(L, int32(idx))
}

func CheckStack(L State, sz int) int {
	return int(lua.checkstack(L, int32(sz)))
}

func XMove(from, to State, n int) {
	lua.xmove(from, to, int32(n))
}

func IsNumber(L State, idx int) bool {
	return lua.isnumber(L, int32(idx)) == 1
}

func IsString(L State, idx int) bool {
	return lua.isstring(L, int32(idx)) == 1
}

func IsCFunction(L State, idx int) bool {
	return lua.iscfunction(L, int32(idx)) == 1
}

func IsUserData(L State, idx int) bool {
	return lua.isuserdata(L, int32(idx)) == 1
}

func Type(L State, idx int) int {
	return int(lua.type_(L, int32(idx)))
}

func TypeName(L State, tp int) string {
	return lua.typename(L, int32(tp))
}

func Equal(L State, idx1, idx2 int) bool {
	return lua.equal(L, int32(idx1), int32(idx2)) == 1
}

func RawEqual(L State, idx1, idx2 int) bool {
	return lua.rawequal(L, int32(idx1), int32(idx2)) == 1
}

func LessThan(L State, idx1, idx2 int) bool {
	return lua.lessthan(L, int32(idx1), int32(idx2)) == 1
}

func ToNumber(L State, idx int) Number {
	return lua.tonumber(L, int32(idx))
}

func ToInteger(L State, idx int) Integer {
	return lua.tointeger(L, int32(idx))
}

func ToBoolean(L State, idx int) bool {
	return lua.toboolean(L, int32(idx))
}

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
	type_       func(L State, idx int32) int32 `lua:"lua_type"`
	typename    func(L State, tp int32) string `lua:"lua_typename"`

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
