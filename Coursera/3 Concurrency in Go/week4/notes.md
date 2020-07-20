## 4.1.1 Iterating through a channel

```go
for i := range c {
    fmt.Println(i)
}
```

Exits loop when channel closes `close(c)`

## 4.1.2 Receiving from multiple channels
 
```go
switch {
    case a := <- c1:
        ...
    case b := <- c2:
        ...
}
```

## 4.1.3 Select with Abort channel

```go
for{
    switch {
        case a := <- c1:
            ...
        case <-abort:
            ...
    }
}
```

## 4.2 Mutexes : Mutual Exclusion

```go
    sync.Mutex
```

## 4.3 Mutex methods
```go
    Lock()
    Unlock()
```

## 4.3.1 Once Sync
f is a function 
```go
    sync.Once.Do(f)
```

### 4.3.2 Deadlocks
Circular dependencies
