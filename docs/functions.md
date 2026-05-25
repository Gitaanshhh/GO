---
title: Functions
---

# Functions

## Signatures and Returns

```go
func add(a, b int) int { return a + b }

func divide(a, b float64) (float64, error) {
  if b == 0 { return 0, errors.New("div by zero") }
  return a / b, nil
}
```

## Variadic and Closures

```go
func sum(nums ...int) int {
  total := 0
  for _, n := range nums {
    total += n
  }
  return total
}

counter := func() func() int {
  n := 0
  return func() int {
    n++
    return n
  }
}
```

## Defer

```go
defer f.Close()
```

`defer` is the standard cleanup mechanism and runs in last-in, first-out order.
