package enum_test

import (
	"fmt"
	"strings"

	"github.com/esome/golang-enum"
)

func ExampleNew() {
	type Color enum.Member[string]

	var (
		Red    = Color{"red"}
		Green  = Color{"green"}
		Blue   = Color{"blue"}
		Colors = enum.New(Red, Green, Blue)
	)

	fmt.Printf("%#v\n", Colors)
	// Output: enum.New(enum_test.Color{"red"}, enum_test.Color{"green"}, enum_test.Color{"blue"})
}

func ExampleEnum_String() {
	type Color enum.Member[string]

	var (
		Red    = Color{"red"}
		Green  = Color{"green"}
		Blue   = Color{"blue"}
		Colors = enum.New(Red, Green, Blue)
	)

	fmt.Println(Colors)
	// Output: red, green, blue
}

func ExampleEnum_GoString() {
	type Color enum.Member[string]

	var (
		Red    = Color{"red"}
		Green  = Color{"green"}
		Blue   = Color{"blue"}
		Colors = enum.New(Red, Green, Blue)
	)

	fmt.Printf("%#v\n", Colors)
	// Output: enum.New(enum_test.Color{"red"}, enum_test.Color{"green"}, enum_test.Color{"blue"})
}

func ExampleEnum_Parse() {
	type Color enum.Member[string]

	var (
		Red    = Color{"red"}
		Green  = Color{"green"}
		Blue   = Color{"blue"}
		Colors = enum.New(Red, Green, Blue)
	)

	parsed := Colors.Parse("red")
	fmt.Printf("%#v\n", parsed)
	// Output: &enum_test.Color{Val:"red"}
}

func ExampleEnum_Contains() {
	type Color enum.Member[string]

	var (
		Red    = Color{"red"}
		Green  = Color{"green"}
		Blue   = Color{"blue"}
		Colors = enum.New(Red, Green, Blue)
	)

	contains := Colors.Contains(Red)
	fmt.Println(contains)
	// Output: true
}

func ExampleEnum_Empty() {
	type Color enum.Member[string]

	var (
		Red    = Color{"red"}
		Green  = Color{"green"}
		Blue   = Color{"blue"}
		Colors = enum.New(Red, Green, Blue)
	)

	empty := Colors.Empty()
	fmt.Println(empty)
	// Output: false
}

func ExampleEnum_Index() {
	type Color enum.Member[string]

	var (
		Red    = Color{"red"}
		Green  = Color{"green"}
		Blue   = Color{"blue"}
		Colors = enum.New(Red, Green, Blue)
	)

	index := Colors.Index(Green)
	fmt.Println(index)
	// Output: 1
}

func ExampleEnum_Value() {
	type Color enum.Member[string]

	var (
		Red    = Color{"red"}
		Green  = Color{"green"}
		Blue   = Color{"blue"}
		Colors = enum.New(Red, Green, Blue)
	)

	value := Colors.Value(Green)
	fmt.Println(value)
	// Output: green
}

func ExampleEnum_Len() {
	type Color enum.Member[string]

	var (
		Red    = Color{"red"}
		Green  = Color{"green"}
		Blue   = Color{"blue"}
		Colors = enum.New(Red, Green, Blue)
	)

	length := Colors.Len()
	fmt.Println(length)
	// Output: 3
}

func ExampleEnum_Members() {
	type Color enum.Member[string]

	var (
		Red    = Color{"red"}
		Green  = Color{"green"}
		Blue   = Color{"blue"}
		Colors = enum.New(Red, Green, Blue)
	)

	members := Colors.Members()
	fmt.Println(members)
	// Output: [{red} {green} {blue}]
}

func ExampleEnum_Values() {
	type Color enum.Member[string]

	var (
		Red    = Color{"red"}
		Green  = Color{"green"}
		Blue   = Color{"blue"}
		Colors = enum.New(Red, Green, Blue)
	)

	values := Colors.Values()
	fmt.Println(values)
	// Output: [red green blue]
}

func ExampleEnum_TypeName() {
	type Color enum.Member[string]

	var (
		Red    = Color{"red"}
		Green  = Color{"green"}
		Blue   = Color{"blue"}
		Colors = enum.New(Red, Green, Blue)
	)

	tname := Colors.TypeName()
	fmt.Println(tname)
	// Output: string
}

func ExampleEnum_Diff() {
	type Color enum.Member[string]

	var (
		Red    = Color{"red"}
		Green  = Color{"green"}
		Blue   = Color{"blue"}
		Purple = Color{"purple"}
		Yellow = Color{"yellow"}
		Colors = enum.New(Red, Green, Blue)
	)

	others := enum.New(Purple, Green, Blue, Yellow)
	// left side
	diffLeft := Colors.Diff(others)
	fmt.Println(diffLeft)
	// right side
	diffRight := others.Diff(Colors)
	fmt.Println(diffRight)
	// Output:
	// red
	// purple, yellow
}

func ExampleEnum_Intersect() {
	type Color enum.Member[string]

	var (
		Red    = Color{"red"}
		Green  = Color{"green"}
		Blue   = Color{"blue"}
		Purple = Color{"purple"}
		Yellow = Color{"yellow"}
		Colors = enum.New(Red, Green, Blue)
	)

	others := enum.New(Purple, Green, Blue, Yellow)
	intersect := Colors.Intersect(others)
	fmt.Println(intersect)
	// Output: green, blue
}

func ExampleEnum_Join() {
	type Color enum.Member[string]

	var (
		Red    = Color{"red"}
		Green  = Color{"green"}
		Blue   = Color{"blue"}
		Purple = Color{"purple"}
		Yellow = Color{"yellow"}
		Colors = enum.New(Red, Green, Blue)
	)

	others := enum.New(Purple, Green, Blue, Yellow)
	joined := Colors.Join(others)
	fmt.Println(joined)
	// Output: red, green, blue, purple, yellow
}

func ExampleNewBuilder() {
	type Color enum.Member[string]
	var (
		b      = enum.NewBuilder[string, Color]()
		Red    = b.Add(Color{"red"})
		Green  = b.Add(Color{"green"})
		Blue   = b.Add(Color{"blue"})
		Colors = b.Enum()
	)

	fmt.Println(
		Colors.Contains(Red),
		Colors.Contains(Green),
		Colors.Contains(Blue),
	)
	// Output:
	// true true true
}

type FoldedString string

// Equal implements [enum.Equaler].
//
// Compare strings ignoring the case.
func (s FoldedString) Equal(other FoldedString) bool {
	return strings.EqualFold(string(s), string(other))
}

func ExampleParse() {
	type Color enum.Member[FoldedString]

	var (
		Red    = Color{"red"}
		Green  = Color{"green"}
		Blue   = Color{"blue"}
		Colors = enum.New(Red, Green, Blue)
	)

	parsed := enum.Parse(Colors, "RED")
	fmt.Printf("%#v\n", parsed)
	// Output: &enum_test.Color{Val:"red"}
}
