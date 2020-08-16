package golin

import "testing"

func TestGetUndefinedValue(t *testing.T) {
    var env = NewEnv()
    var symbol string = "foo"
    var result = env.Get(symbol)
    
    if result != nil {
        t.Fail()
    }
}

func TestSetAndFindValue(t *testing.T) {
    var env = NewEnv()
    var symbol string = "foo"
    var atom = Atom{"bar"}
    var _ SExpr = (*Atom)(nil)
    env.Set(symbol, atom)
    var result = env.Get(symbol)

    if result.(Atom).name != "bar" {
        t.Fail()
    }
}

