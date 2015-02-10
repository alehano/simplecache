package simplecache

import (
    "time"
)

/*
Usage:

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

*/


type Cache struct {
    expire map[string]time.Time
}

func (c *Cache) Init() {
    c.expire = make(map[string]time.Time)
}

func (c Cache) IsValid(name string) bool {
    if t, ok := c.expire[name]; ok && t.After(time.Now()) {
        return true
    }
    return false
}

func (c *Cache) SetExpireTime(name string, time time.Time) {
    c.expire[name] = time
}

func (c *Cache) SetExpireDur(name string, dur time.Duration) {
    c.SetExpireTime(name, time.Now().Add(dur))
}

func (c *Cache) Invalidate(name string) {
    c.SetExpireTime(name, time.Now())
}


