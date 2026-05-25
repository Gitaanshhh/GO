---
title: Arrays and Slices
---

# Arrays and Slices

## Arrays

```go
var a [3]int
b := [3]int{1, 2, 3}
c := [...]int{1, 2, 3}

x := b
x[0] = 99
```

Arrays have fixed size and copy by value.

## Slices

```go
s := []int{1, 2, 3}
s2 := make([]int, 5)
s3 := make([]int, 3, 5)

s = append(s, 4, 5)

sub := s[1:4]
dst := make([]int, len(s))
copy(dst, s)
```

Slices are dynamic views over an array, so sub-slices share backing storage unless you copy them.

## Range

```go
for i, v := range s {
  _ = i
  _ = v
}

matrix := [][]int{[]int{1, 2}, []int{3, 4}}
```
