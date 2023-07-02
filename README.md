# Let's Go workshop

## Introduction

This workshop aims to give an introduction of the Go language by writing code and practising the language. Step by step more code will be introduced and the final product will be a tool that creates GIFs from an input text.

## Tools needed for the workshop

- Go
  - Download at go.dev
  - Verify that it's working by running `go version` in the command line
- The IDE of your choice
  - Visual Studio Code with the Go extension
  - Goland (from JetBrains)
  - Vim plugin <https://github.com/fatih/vim-go>
- Git
  - To commit the different milestones
- pkg.go.dev
  - This is the link to the documentation of the various Go packages available
    - Inside or outside the standard library
  - Useful to learn about each symbol (types, functions, methods, and so on) and what they do for each package. It usually contains examples of usages of these symbols as well

## Milestones

Here is the list of different milestones we are going to acheive during the workshop.

### Hello, World! (milestone 0)

For the first step you are going to write the hello world, and this is also going to be the occasion to setup your Go project.

#### Steps

1. Create a folder, e.g. lets-go-workshop; This will contain the code to write for the workshop
2. Go inside the folder and run `go mod init` from the terminal
3. (optional) run `git init` if you want to store the code for the different milestones in a different commit
4. Open your IDE inside the workshop folder
5. Create a `main.go` file
6. Write the following code:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

7. Run `go run main.go`

#### Explanations

The command `go mod init` initializes a project making it a Go __module__. A Go module represents a set of Go files that are supposed to be shipped altogether: e.g. an application or a library.

Every Go application starts at a `main()` function located in a `main.go` file.

Every Go file resides in a `package`, this information is declared as the first instruction of the file. The main function should reside in __package main__.

Packages can contain functions, constants, custom types and so on inside a specific namespace; to use a specific package it needs to be imported with the `import` keyword. An unused import results in a compilation error.

Any exported/public symbol inside a package can be invoked in the form of `package name.symbol name`. For example, in the call `fmt.Println`, __fmt__ is the package name and __Println__ is the symbol name.

In Go `;` is not mandatory at the end of each statement.

### Milestone 1

For this step you are going to improve on the hello world with some code that creates a file and writes the hello world on it. The goal is to see __variable__ declaration and initialization, __conditionals__ with the __if__ statement, to start interacting with some packages the __standard library__ and to introduce few Go specific features: the __blank identifier__, the __multi-value returns__, the __error__ type and the __defer__ keyword.

#### Steps for milestone 1

1. Create a file (use the `os` package), write the "hello, world!" string to the file (use the `fmt` package), close the file
2. if there are any errors during the file creation, log them and exit (use the `log` package)
3. Use the defer keyword to defer the execution of the Close() call
4. Use the command line arguments to take the text to print (use the `flag` package)
5. Run `go run main.go "Hello, World!"`

#### Notes for milestone 1

Use [pkg.go.dev](pkg.go.dev) to search the content of the packages `os`, `fmt`, `log` and `flag`.

To declare and assign __variables__ you can use two tools:

- The `:=` operator (declare and assign)
  - a := 1
- The `var` keyword (declare), followed potentially by the `=` operator (assign)
  - var a int = 1

When the compiler can infer the type of a variable it can be omitted:

- var a int = 1 // => var a = 1
- a := 1
  - in this statement the type is already omitted

When the type name is present in a statement, it __always__ goes after the name:

- var a int // :+1:
- var int a // :x: (compilation error)

You can declare and/or assign multiple variables, of the same or different type, at the same time:

- a, b, c := "hello", 1, "world"
- f, err := os.Create("myFile") // os.Create is a function that returns 2 values

If you don't need a specific value during an assignment you can ignore it using the __blank identifier__ `_`. Any __unused variables__ are considered complier errors.

- f, _ := os.Create("myFile") // following the previous example, I just need the value of f from the function

The __if__ statement is similar to most languages, here are some differences:

- No `()` are required around the condition
- `{}` are always required around the code to execute if the condition is true, even if it fits in one line

The __defer__ keyword instructs the function that follows to be called at the end of the function where it is deferred.

Generally, in a __multi-value return__ the second argument or last argument is of type __error__ which is an interface that declares one method: `func Error() string`

A value of a specific interface type can be compared only with another value of the same type or with `nil` which means: no value for that interface.

### Milestone 2

In this milestone you'll create your first image (a PNG); in doing so you'll learn about creating your own __package__ and __function__, invoke them and learn about __for__ loops, __types__, __zero values__ and the meaning of __exported__ and __unexported__.

#### Steps for milestone 2

1. Create a new folder named __myimage__; in Go, this is a new package
2. Inside the myimage folder create a new file called __myimage.go__
3. Write a function `func New(l, h int, c color.RGBA) *image.Paletted` inside __myimage.go__ that draws a rectangle of the specified size and color (use the `image` package, see functions `Rect`, `NewPaletted` and the `Set` method of the `Paletted` type)
4. Change the main.go to first call the function that draws the rectangle then use `png.Encode` to save the rectangle image in a PNG file
5. Run `go run main.go`

As a bonus you can add more input flags for the properties of the rectangle (length, height and color) and the output file name.

#### Notes for milestone 2

A __package name__ and its containing folder do not have to have matching names, but it's good practice to do so.

There are no specific keywords to define the scope of a Go symbol:

- Name is UPPERCASE
  - Symbol is __exported__, which means available to all packages (public-level scope)
- Name is lowercase
  - Symbol is __unexported__, which means available to the package where it is located (package-level scope)

There three kind of __types__ in Go:

1. Primitives
2. Internal types
3. Custom types

Primitives in Go are __bool__ (true or false), the literal type __string__ and __numeric__ types __rune__, __byte__ (rune and byte can also represent literals), __uint__, uint8, uint16, uint32, uint64, __int__, int8, int16, int32, int64, float32, __float64__, complex64, __complex128__).

Internal types are array, function, slice, map, and channel.

Custom types are all Go types declared with the `type` and/or `struct` keywords.

For example:

```go
type Age int

type Person struct {
  name string
  age  Age
}
```

Every variable is initialized in Go at the declaration time; if not explicitly initialized in the code, they will be initialized by Go with the __zero value__ for the type of that variable.

Here is the zero value list:

```go
var (
  b bool // false
  i int // 0 (same for all numeric types)
  s string // ""
  p Person // Person{}, all custom types are initialized with all fields set to their zero values
  v map[int]string // nil, also valid for all internal types
)
```

Looping in Go is similar to most languages with some differences:

- `for` is the only keyword used for loops
- No `()` are required around the loop instructions
- `{}` are always required around the code block to execute inside the loop, even if it fits in one line
- __for condition {}__ replaces the _while_ statement; there are no _do/while_ statements in Go
- __for true {}__ represents an infinite loop (to be interrupted with a `break` statement

Using flags in Go requires the usage of the `flag` package:

- first define all flags using the functions `Int`, `String`, `Bool` and so on, according to the type of the data you request in input
  - They all take __flag name__, __default value__ and __short description__ as inputs and return pointers of the same type as in the function name
  - They can be visualized when running `go run main.go --help`, without needing to define a flag for the help message
- then call the function `flag.Parse()` to parse all flags defined above

__Pointers__ in Go are similar to C/C++. Use `*` declare or dereference them, use `&` to take the address of a value.

### Milestone 3

In this milestone you'll write your first message on the previously created image; in doing so you'll learn interacting with __external packages__ from the standard library and see how to import them in your project.

#### Steps for milestone 3

1. Write a function `func Write(dst draw.Image, text string, c color.RGBA, fontPath string, fontSize float64) error` inside __myimage.go__ that draws a text on the rectangle created in milestone 2
2. Import the package `github.com/golang/freetype` inside __myimage.go__
3. In the terminal, at the root of the project run `go mod tidy`
4. Write the code to create and configure the `freetype.Context` object
5. Call the `Write` function inside the main function,
6. Run `go run main.go`

As a bonus you can:

- keep adding input flags for the color of the string, the font file path and the font size.
- create a function that gives the appropriate size for the rectangle by taking into account the font size and the length of the text to draw

#### Notes for milestone 3

The difference between packages from the __standard library__ and the __external__ ones can be seen from its naming when getting or importing them. In fact, the main difference is that external packages include a complete path to external package containing:

- the registry: github, gitlab, bitbucket, and so on
- the package author: be it a username or a project group name
- the package name and path to subpackage if relevant
  - in local it will be the path to the package

the function `Join` of the `strings` package is a useful tool to join several strings, organized into a slice (array), in a single string.
