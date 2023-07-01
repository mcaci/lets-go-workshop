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
- `{}` are always required adound the code to execute if the condition is true, even if it fits in one line

The __defer__ keyword instructs the function that follows to be called at the end of the function where it is deferred.

Generally, in a __multi-value return__ the second argument or last argument is of type __error__ which is an interface that declares one method: `func Error() string`

A value of a specific interface type can be compared only with another value of the same type or with `nil` which means: no value for that interface.
