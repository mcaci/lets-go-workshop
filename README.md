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
