# Go Notes

This repository now keeps the long-form notes in a GitHub Pages-friendly `docs/` folder, with one markdown page per topic and a landing page for quick navigation.

Use [docs/index.md](docs/index.md) as the entry point for the published site. The topic pages mirror the cheat-sheet structure from the attached reference and are grouped by concept so they are easy to browse, expand, and maintain.

## Docs Map

* [Basics](docs/basics.md)
* [Arrays & Slices](docs/arrays-and-slices.md)
* [Maps](docs/maps.md)
* [Structs](docs/structs.md)
* [Control Flow](docs/control-flow.md)
* [Functions](docs/functions.md)
* [Packages & Scope](docs/packages-and-scope.md)
* [Concurrency](docs/concurrency.md)
* [Errors](docs/errors.md)
* [Hosting GitHub Pages](docs/hosting.md)

## Why Go?

Go was designed for:

* multicore systems
* cloud infrastructure
* distributed systems
* backend/network services

Key advantages:

* built-in concurrency
* lightweight threading via goroutines
* fast compile + execution
* simple syntax
* statically typed
* compiled to a single binary

Combines:

* Python-like simplicity
* C/C++-like performance

---

# Project Structure

## Initialize Module

```bash
go mod init myapp
```

Creates:

```text
go.mod
```

Contains:

* module name/path
* Go version
* dependencies

---

# Packages

Every file starts with:

```go
package main
```

Rules:

* one package per folder
* files in same folder usually same package
* only `package main` + `func main()` creates executable

---

## Structure Hierarchy

```text
Module  -> entire project
Package -> folder of related files
File    -> source file
```

---

# Running Code

## Run

```bash
go run .
```

or:

```bash
go run main.go
```

`go run .`:

* compiles temporarily
* runs immediately
* no executable generated

---

## Build Executable

```bash
go build .
```

Runs generated binary:

```bash
./projectname
```

Windows:

```bash
.\projectname.exe
```

---

# Variables & Constants

## Declaration

```go
var x int = 5
var y = 5
z := 5              // funcs only

var a, b int = 1, 2
```

---

## Constants

```go
const Pi = 3.14
```

Grouped:

```go
const (
    A = 1
    B = 2
)
```

---

## iota

Auto-increment enum generator.

```go
const (
    Red = iota   // 0
    Green        // 1
    Blue         // 2
)
```

---

# Types

## Primitive Types

```go
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64

float32 float64

bool
string
byte   // alias for uint8
rune   // alias for int32
```

---

## Zero Values

Unlike C, uninitialized vars are safe.

```go
var i int       // 0
var f float64   // 0.0
var b bool      // false
var s string    // ""
var p *int      // nil
```

---

## Type Conversion

No implicit conversion.

```go
x := 5
y := float64(x)
```

---

# Pointers

```go
x := 10

p := &x
fmt.Println(*p)

*p = 20
fmt.Println(x)
```

---

## Notes

* no pointer arithmetic
* garbage collected
* structs auto-deref

```go
p.Name
```

instead of:

```c
p->Name
```

---

# Arrays

Fixed size.

```go
var arr [3]int
```

Initialization:

```go
arr := [3]int{1, 2, 3}
```

Compiler counts:

```go
arr := [...]int{1, 2, 3}
```

---

## Arrays are Value Types

```go
a := [3]int{1,2,3}
b := a

b[0] = 99

fmt.Println(a) // unchanged
```

---

# Slices

Dynamic arrays.

```go
nums := []int{1,2,3}
```

Using `make`:

```go
s := make([]int, 5)
s2 := make([]int, 3, 10)
```

```text
make(type, len, cap)
```

---

## Append

```go
nums = append(nums, 4)
nums = append(nums, 5, 6)
```

Spread operator:

```go
nums = append(nums, otherSlice...)
```

---

## Slice Operations

```go
s[1:4]
s[:4]
s[2:]
```

---

## len & cap

```go
len(s)
cap(s)
```

---

## copy

```go
dst := make([]int, len(src))
copy(dst, src)
```

Needed because sub-slices share memory.

---

# Loops

Only `for` exists.

---

## C-style

```go
for i := 0; i < 5; i++ {
}
```

---

## While-style

```go
for x < 100 {
    x *= 2
}
```

---

## Infinite Loop

```go
for {
}
```

---

## range

```go
for i, v := range nums {
    fmt.Println(i, v)
}
```

Ignore value:

```go
for i := range nums {
}
```

Ignore index:

```go
for _, v := range nums {
}
```

---

# Conditionals

## if

```go
if x > 5 {

} else if x == 5 {

} else {

}
```

---

## if with init

Very common.

```go
if err := doThing(); err != nil {
    fmt.Println(err)
}
```

`err` scoped only to block.

---

# Switch

No automatic fallthrough.

```go
switch x {
case 1:
    fmt.Println("one")
case 2,3:
    fmt.Println("two or three")
default:
    fmt.Println("other")
}
```

---

## Expressionless Switch

```go
switch {
case x < 0:
case x == 0:
case x > 0:
}
```

---

## fallthrough

```go
switch x {
case 1:
    fmt.Println("one")
    fallthrough
case 2:
    fmt.Println("two")
}
```

---

## Type Switch

```go
switch v := any.(type) {
case int:
case string:
}
```

---

# Functions

## Basic

```go
func add(a int, b int) int {
    return a + b
}
```

Same-type shorthand:

```go
func add(a, b int) int {
    return a + b
}
```

---

## Multiple Returns

Idiomatic Go.

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("divide by zero")
    }

    return a / b, nil
}
```

Usage:

```go
result, err := divide(10, 2)
```

---

## Named Returns

```go
func f() (x int, y int) {
    x = 1
    y = 2
    return
}
```

---

## Variadic Functions

```go
func sum(nums ...int) int {
    total := 0

    for _, n := range nums {
        total += n
    }

    return total
}
```

Pass slice:

```go
sum(slice...)
```

---

# First-Class Functions

## Anonymous Functions

```go
double := func(x int) int {
    return x * 2
}
```

---

## Closures

```go
func counter() func() int {
    n := 0

    return func() int {
        n++
        return n
    }
}
```

---

# defer

Runs when function exits.

```go
func readFile() {
    f, _ := os.Open("a.txt")
    defer f.Close()
}
```

Multiple defers = LIFO.

---

# Maps

## Creation

```go
m := map[string]int{
    "a": 1,
    "b": 2,
}
```

Using make:

```go
m := make(map[string]int)
```

---

## Operations

```go
m["x"] = 10

v := m["x"]

delete(m, "x")
```

---

## Existence Check

```go
v, ok := m["x"]
```

`ok == false` if absent.

---

## Iteration

```go
for k, v := range m {
}
```

Map order is NOT guaranteed.

---

# Structs

## Definition

```go
type Person struct {
    Name string
    Age  int
}
```

---

## Initialization

```go
p := Person{
    Name: "Alice",
    Age: 20,
}
```

Positional:

```go
p := Person{"Alice", 20}
```

---

# Methods

## Value Receiver

Gets copy.

```go
func (p Person) Greet() {
    fmt.Println(p.Name)
}
```

---

## Pointer Receiver

Can modify original.

```go
func (p *Person) Birthday() {
    p.Age++
}
```

---

# Embedding

Composition instead of inheritance.

```go
type Employee struct {
    Person
    Company string
}
```

Access directly:

```go
e.Name
e.Greet()
```

---

# Interfaces

## Definition

```go
type Greeter interface {
    Greet()
}
```

---

## Implicit Implementation

No `implements` keyword.

```go
type Person struct {}

func (p Person) Greet() {}

var g Greeter = Person{}
```

---

## Empty Interface

Can hold anything.

```go
var x interface{}
```

Equivalent idea to Java `Object`.

---

## Type Assertion

```go
s, ok := x.(string)
```

---

# Packages & Imports

## Imports

```go
import "fmt"
```

Grouped:

```go
import (
    "fmt"
    "math"
)
```

---

## Alias Import

```go
import m "math"
```

Usage:

```go
m.Sqrt(16)
```

---

## Blank Import

```go
import _ "image/png"
```

Imports only for side effects.

---

# Visibility Rules

Capitalized = exported/public.

```go
func PublicFunc() {}
```

Lowercase = package-private.

```go
func privateFunc() {}
```

Same for structs/fields.

---

# Scope Rules

## Block Scope

```go
x := 1

{
    x := 2
}
```

Inner `x` shadows outer.

---

## Package Scope

Declared outside funcs.

```go
var Global int
```

---

## init()

Runs before `main`.

```go
func init() {
}
```

No args or returns.

---

# Error Handling

Errors are values, not exceptions.

---

## Basic Pattern

```go
val, err := strconv.Atoi("42")

if err != nil {
    return err
}
```

---

## Creating Errors

```go
errors.New("bad input")
```

---

## Wrapping Errors

```go
fmt.Errorf("context: %w", err)
```

---

## errors.Is / errors.As

```go
errors.Is(err, target)
errors.As(err, &target)
```

---

# panic & recover

## panic

```go
panic("fatal")
```

Avoid for normal flow.

---

## recover

Works only inside deferred funcs.

```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println(r)
    }
}()
```

---

# Concurrency

# Goroutines

Lightweight threads managed by Go runtime.

```go
go myFunc()
```

Anonymous:

```go
go func() {
    fmt.Println("hello")
}()
```

---

# WaitGroup

Wait for goroutines.

```go
var wg sync.WaitGroup

wg.Add(1)

go func() {
    defer wg.Done()
}()

wg.Wait()
```

---

# Channels

## Create

```go
ch := make(chan int)
```

Buffered:

```go
ch := make(chan int, 5)
```

---

## Send / Receive

```go
ch <- 10

x := <-ch
```

---

## Close

```go
close(ch)
```

---

## Range Over Channel

```go
for v := range ch {
}
```

---

## Directional Channels

```go
func producer(out chan<- int) {}
func consumer(in <-chan int) {}
```

---

# select

Like switch for channels.

```go
select {
case v := <-ch1:
    fmt.Println(v)

case ch2 <- 5:

default:
}
```

---

# Mutex

Protect shared memory.

```go
var mu sync.Mutex

mu.Lock()
shared++
mu.Unlock()
```

Usually:

```go
mu.Lock()
defer mu.Unlock()
```

---

# Common Built-ins

```go
len()
cap()
append()
copy()
make()
new()
delete()
close()
panic()
recover()
```

---

# Common Commands

## Create Module

```bash
go mod init myapp
```

---

## Run

```bash
go run .
```

---

## Build

```bash
go build .
```

---

## Format

```bash
go fmt
```

---

## Fetch/Clean Dependencies

```bash
go mod tidy
```

---

# Important Go Differences

```text
No classes
No inheritance
No exceptions
No while/do-while
No ternary operator
No function overloading
No implicit conversions
No pointer arithmetic
```

---

# Go Idioms

## Error Handling

```go
if err != nil {
    return err
}
```

---

## Short Declaration

```go
x := 5
```

---

## defer Cleanup

```go
defer file.Close()
```

---

## Composition Over Inheritance

Use embedding.

---

# Generics (Go 1.18+)

```go
func Min[T comparable](a, b T) T {
    if a == b {
        return a
    }

    return b
}
```

---

# Minimal Program

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

