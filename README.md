# Go LuaJIT

[CGO-less](https://github.com/ebitengine/purego) bindings *(+ optional Go-like wrapper)* for [LuaJit](https://luajit.org/).

## Usage

```sh
go get github.com/judah-caruso/go-luajit@latest
```

Copy the appropriate libraries from this directory into your project:

- `libluajit.dll` (Windows)
- `libluajit.dylib` (Mac)
- `libluajit.so` (Linux)

## Wrapper

Aside from the 1:1 bindings (`go-luajit/lua`), an optional Go-like wrapper is provided by importing `go-luajit`. This wrapper makes LuaJIT easier to use from the Go, removing some of the underlying C-isms from the library.

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
