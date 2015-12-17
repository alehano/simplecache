# Simple Cache

Simple wrapper for any struct that add cache functionallity.

Methods: 
```
IsValid(string) bool
SetExpireDur(string, time.Duration)
SetExpireTime(string, time.Time)
Invalidate(string)
```

Usage:

```go
// Set your own cache type
type MyType int
type MyCache struct {
    Cache
    data map[string]MyType
}

// Init it
mc := &MyCache{data: make(map[string]MyType)}
mc.Init()

// Use it
if(mc.IsValid("key")) {
    fmt.Println(mc.data["key"])
} else {
    // compute value
    var val MyType
    val = 1
    mc.SetExpireDur("key", time.Second * 60)
    mc.data["key"] = val
    fmt.Println(mc.data["key"])
}

// Invalidate
mc.Invalidate("key")

```
