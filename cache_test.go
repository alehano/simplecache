package simplecache

import (
    "testing"
    "time"
)

func TestCache(t *testing.T) {

    type MyType int

    type MyCache struct {
        Cache
        data map[string]MyType
    }

    mc := &MyCache{}
    mc.Init()
    mc.data = make(map[string]MyType)

    if(mc.IsValid("key")) {
        t.Error("Shouldn't be set")
    } else {
        // compute value
        var val MyType
        val = 1
        mc.SetExpireDur("key", time.Second * 3)
        mc.data["key"] = val
    }

    if(!mc.IsValid("key")) {
        t.Error("Should be set")
    }

    time.Sleep(time.Second * 3)

    if(mc.IsValid("key")) {
        t.Error("Shouldn't be set")
    }

    mc.SetExpireDur("key", time.Second * 60)
    mc.Invalidate("key")

    if(mc.IsValid("key")) {
        t.Error("Shouldn't be set")
    }

    mc.SetExpireTime("key", time.Now())

    if(mc.IsValid("key")) {
        t.Error("Shouldn't be set")
    }

    mc.SetExpireTime("key", time.Now().Add(time.Minute * -1))

    if(mc.IsValid("key")) {
        t.Error("Shouldn't be set")
    }

}
