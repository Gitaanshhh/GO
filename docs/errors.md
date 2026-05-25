---
title: Errors
---

# Errors

## Error Handling

```go
val, err := strconv.Atoi("42")
if err != nil {
  return fmt.Errorf("parse failed: %w", err)
}

err1 := errors.New("base error")
err2 := fmt.Errorf("context: %w", err1)
```

## Custom Errors and Panic

```go
type AppError struct {
  Code int
  Msg  string
}

func (e *AppError) Error() string {
  return fmt.Sprintf("[%d] %s", e.Code, e.Msg)
}

panic("unrecoverable!")
```

## Recover

```go
defer func() {
  if r := recover(); r != nil {
    fmt.Println("recovered:", r)
  }
}()
```

Use panic and recover only for unexpected failures, not as regular control flow.
