package golin

import "testing"

func TestFindUndefinedValue(t *testing.T) {
    var env = Env{}
    var symbol Symbol = "foo"
    var result = env.Find(symbol)
    
    if result != nil {
        t.Fail()
    }
}