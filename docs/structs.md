---
title: Structs
---

# Structs

## Structs and Methods

```go
type Person struct {
  Name string
  age  int
}

func (p Person) Greet() string {
  return "Hi, " + p.Name
}

func (p *Person) Birthday() {
  p.age++
}
```

## Embedding

```go
type Employee struct {
  Person
  Company string
}

e := Employee{
  Person:  Person{Name: "Ali", age: 25},
  Company: "Cisco",
}
```

## Interfaces

```go
type Greeter interface {
  Greet() string
}

var g Greeter = Person{Name: "Alice"}
```

Go uses structural typing, so a type satisfies an interface implicitly when it has the right methods.
