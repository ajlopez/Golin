package golin

import "testing"

func TestFindUndefinedValue(t *testing.T) {
    var env = NewEnv()
    var symbol Symbol = "foo"
    var result = env.Find(symbol)
    
    if result != nil {
        t.Fail()
    }
}

func TestSetAndFindValue(t *testing.T) {
    var env = NewEnv()
    var symbol Symbol = "foo"
    env.Set(symbol, 42)
    var result = env.Find(symbol)
    
    if result != 42 {
        t.Fail()
    }
}