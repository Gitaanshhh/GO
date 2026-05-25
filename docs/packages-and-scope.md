---
title: Packages and Scope
---

# Packages and Scope

## Packages and Imports

```go
package main

import (
  "fmt"
  m "math"
  _ "image/png"
)
```

## Visibility

```go
type PublicStruct struct {}
type privateStruct struct {}

func PublicFunc() {}
func privateFunc() {}
```

Capitalized names are exported, lowercase names stay package-private.

## init

```go
func init() {}
```

`init()` runs before `main()` and is useful for package setup.
