---
title: Basics
---

# Basics

## Variables and Constants

```go
var x int = 5
var y = 5
z := 5

var a, b int = 1, 2

const Pi = 3.14

const (
  A = iota
  B
  C
)
```

## Types and Zero Values

```go
var i int
var f float64
var b bool
var s string
var p *int
```

## Pointers

```go
x := 42
p := &x
*p = 100

ps := &Person{Name: "Ali"}
ps.Name

p2 := new(int)
```

`:=` is the usual short declaration form, and Go requires unused variables to be removed before compilation.
