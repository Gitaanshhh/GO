---
title: Maps
---

# Maps

## Map Basics

```go
m := map[string]int{"a": 1, "b": 2}
m2 := make(map[string]int)

m["key"] = 42
val := m["key"]
val, ok := m["key"]
delete(m, "key")
```

Use the two-value lookup when you need to distinguish a missing key from a zero value.

## Map of Slices

```go
graph := make(map[string][]string)
graph["a"] = append(graph["a"], "b", "c")
```
