---
title: Control Flow
---

# Control Flow

## For

```go
for i := 0; i < 5; i++ {}

for x < 100 {
  x *= 2
}

for {
  break
}
```

## If and Switch

```go
if err := doThing(); err != nil {
  log.Fatal(err)
}

switch x {
case 1:
  fmt.Println("one")
case 2, 3:
  fmt.Println("two|three")
default:
}
```

Use `fallthrough` explicitly when you want C-style switch behavior.
