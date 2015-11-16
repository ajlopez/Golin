package golin

import "testing"

func TestGetUndefinedValue(t *testing.T) {
    var env = NewEnv()
    var symbol Symbol = "foo"
    var result = env.Get(symbol)
    
    if result != nil {
        t.Fail()
    }
}

func TestSetAndFindValue(t *testing.T) {
    var env = NewEnv()
    var symbol Symbol = "foo"
    env.Set(symbol, 42)
    var result = env.Get(symbol)
    
    if result != 42 {
        t.Fail()
    }
}