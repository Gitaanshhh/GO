---
title: Concurrency
---

# Concurrency

## Goroutines

```go
go func() {
  fmt.Println("concurrent!")
}()

var wg sync.WaitGroup
for i := 0; i < 5; i++ {
  wg.Add(1)
  go func(n int) {
    defer wg.Done()
    work(n)
  }(i)
}
wg.Wait()
```

## Channels and Select

```go
ch := make(chan int)
bch := make(chan int, 5)

ch <- 42
v := <-ch
close(ch)

select {
case v := <-ch1:
  fmt.Println(v)
case <-time.After(time.Second):
  fmt.Println("timeout")
default:
}
```

## Mutex

```go
var mu sync.Mutex
mu.Lock()
sharedCounter++
mu.Unlock()
```
