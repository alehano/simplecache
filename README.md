# Simple Cache
Simple in-memory type safe cache in Go (golang)


Usage:

```go
// Set your own cache type
type MyType int
type MyCache struct {
    Cache
    data map[string]MyType
}

// Init it
mc := &MyCache{}
mc.Init()
mc.data = make(map[string]MyType)

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
