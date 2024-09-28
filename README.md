# Go LuaJIT

[CGO-less](https://github.com/ebitengine/purego) bindings *(+ optional go-like wrapper)* for [LuaJit](https://luajit.org/).

## Usage

```sh
go get github.com/judah-caruso/go-luajit@latest
```

Copy the appropriate libraries from this directory into your project:

- `libluajit.dll` (Windows)
- `libluajit.dylib` (Mac)
- `libluajit.so` (Linux)

## Examples

```go
import "github.com/judah-caruso/go-luajit/lua"

func main() {
	L := lua.NewState()
	defer lua.Close(L)

	lua.OpenLibs(L)
	lua.LoadString(L, `print("hello from luajit!")`)
	lua.Call(L, 0, 0)
}
```

See `examples/` for more examples.

## License

Public domain. See [LICENSE](./LICENSE) for more information.
