package golin2

import "testing"

func TestIntegerValueToString(t *testing.T) {
    var value = Value{42}
    
    var result = value.String()
    
    if result != "42" {
        t.Fatal("value.String() is not \"42\"")
    }
}

func TestEvaluateIntegerValue(t *testing.T) {
    var value = Value{42}
    var env = NewEnv()
    
    var result = value.Evaluate(env)
    
    if result.(Value).value != 42 {
        t.Fatal("value.Evaluate(42) is not Value{42}")
    }
}

