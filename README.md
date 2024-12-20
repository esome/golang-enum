# enum

[ [📄 docs](https://pkg.go.dev/github.com/esome/golang-enum) ] [ [🐙 github](https://github.com/esome/golang-enum) ]

Type safe enums for Go without code generation or reflection.

😎 Features:

* Type-safe, thanks to generics.
* No code generation.
* No reflection.
* Well-documented, with working examples for every function.
* Flexible, supports both static and runtime definitions.
* Zero-dependency.


🧬 Origin: 

This project is a fork of [🐙 https://github.com/orsinium-labs/enum](https://github.com/orsinium-labs/enum).  
So if you like it, consider [❤️ sponsoring](https://github.com/sponsors/orsinium) the original author.

⛓️‍💥️ Breaking changes compared to the original library:

The generic value field of `enum.Member` has been renamed to `Val` from `Value`.
This enables the implementation of the `database/sql/driver.Valuer` interface on derived types.

Unfortunately, `@orsinium` refused to merge the pull-request, so it is a separate fork now.

## 1.1. 📦 Installation

```bash
go get github.com/esome/golang-enum
```

## 1.2. 🛠️ Usage

Define:

```go
type Color enum.Member[string]

var (
  Red    = Color{"red"}
  Green  = Color{"green"}
  Blue   = Color{"blue"}
  Colors = enum.New(Red, Green, Blue)
)
```

Parse a raw value (`nil` is returned for invalid value):

```go
parsed := Colors.Parse("red")
```

Compare enum members:

```go
parsed == Red
Red != Green
```

Accept enum members as function arguments:

```go
func SetPixel(x, i int, c Color)
```

Loop over all enum members:

```go
for _, color := range Colors.Members() {
  // ...
}
```

Ensure that the enum member belongs to an enum (can be useful for defensive programming to ensure that the caller doesn't construct an enum member manually):

```go
func f(color Color) {
  if !colors.Contains(color) {
    panic("invalid color")
  }
  // ...
}
```

Define custom methods on enum members:

```go
// UnmarshalJSON implements the [encoding/json.Unmarshaler] interface
func (c Color) UnmarshalJSON(b []byte) error {
  return nil
}

// Value implements the [database/sql/driver.Valuer] interface
func (c Color) Value() (driver.Value, error) {
  return nil, nil
}
```

Dynamically create enums to pass multiple members in a function:

```go
func SetPixel2(x, y int, colors enum.Enum[Color, string]) {
  if colors.Contains(Red) {
    // ...
  }
}

purple := enum.New(Red, Blue)
SetPixel2(0, 0, purple)
```

Enum members can be any comparable type, not just strings:

```go
type ColorValue struct {
  UI string
  DB int
}
type Color enum.Member[ColorValue]
var (
  Red    = Color{ColorValue{"red", 1}}
  Green  = Color{ColorValue{"green", 2}}
  Blue   = Color{ColorValue{"blue", 3}}
  Colors = enum.New(Red, Green, Blue)
)

fmt.Println(Red.Value.UI)
```

If the enum has lots of members and new ones may be added over time, it's easy to forget to register all members in the enum. To prevent this, use enum.Builder to define an enum:

```go
type Color enum.Member[string]

var (
  b      = enum.NewBuilder[string, Color]()
  Red    = b.Add(Color{"red"})
  Green  = b.Add(Color{"green"})
  Blue   = b.Add(Color{"blue"})
  Colors = b.Enum()
)
```

## 1.3. 🤔 QnA

1. **What happens when enums are added in Go itself?** I'll keep it alive until someone uses it, but I expect the project popularity to quickly die out when there is native language support for enums. When you can mess with the compiler itself, you can do more. For example, this package can't provide an exhaustiveness check for switch statements using enums (maybe only by implementing a linter) but proper language-level enums would most likely have it.
2. **Is it reliable?** Yes, pretty much. It has good tests but most importantly it's a small project with just a bit of the actual code that is hard to mess up.
3. **Is it maintained?** The project will review and merge upstream changes regularly, if they can be adopted easily. However, active development is unlikely to happen, since it is considered feature-complete, very much like the original one.
4. **What if I found a bug?** 
   1. Please check, whether the bug exists in the original project as well, and follow their contribution policies. Then the point above applies.
   2. If it only exists here, fork the project, fix the bug, write some tests, and open a Pull Request. It will be reviewed contemporary.
