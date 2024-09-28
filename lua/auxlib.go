package lua

type LReg struct {
	Name string
	Func CFunction
}

const (
	ErrFile = ErrErr + 1
)

func NewState() State {
	return luaL.newstate()
}

func OpenLib(L State, libname string, l []LReg) {
	luaL.openlib(L, libname, &l[0], int32(len(l)))
}

func Register(L State, libname string, l []LReg) {
	if len(l) != 0 && l[len(l)-1].Func != nil {
		l = append(l, LReg{Func: nil})
	}

	luaL.register(L, libname, &l[0])
}

func GetMetaField(L State, obj int, e string) int {
	return int(luaL.getmetafield(L, int32(obj), e))
}

func CallMeta(L State, obj int, e string) int {
	return int(luaL.callmeta(L, int32(obj), e))
}

func TypeError(L State, narg int, tname string) int {
	return int(luaL.typerror(L, int32(narg), tname))
}

func ArgError(L State, numarg int, extramsg string) int {
	return int(luaL.argerror(L, int32(numarg), extramsg))
}

func LoadString(L State, s string) int {
	return int(luaL.loadstring(L, s))
}

// Macro conversions

func ArgCheck(L State, cond bool, numarg int, extramsg string) bool {
	return cond || luaL.argerror(L, int32(numarg), extramsg) == 1
}

func CheckString(L State, numarg int) string {
	return luaL.checklstring(L, int32(numarg), nil)
}

func OptString(L State, numarg int, def string) string {
	return luaL.optlstring(L, int32(numarg), def, nil)
}

func CheckInt(L State, numarg int) Integer {
	return luaL.checkinteger(L, int32(numarg))
}

func OptInt(L State, numarg int, def Integer) Integer {
	return luaL.optinteger(L, int32(numarg), def)
}

func TypeNameOf(L State, idx int) string {
	return TypeName(L, Type(L, idx))
}

func GetMetaTableFor(L State, name string) {
	GetField(L, RegistryIndex, name)
}

var luaL struct {
	_ nocopy

	openlib      func(L State, libname string, l *LReg, nup int32)  `lua:"luaL_openlib"`
	register     func(L State, libname string, l *LReg)             `lua:"luaL_register"`
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
