package golin

import "fmt"

type Value struct {
    value interface{}
}

func (value Value) Evaluate(env *Env) SExpr {
    return value
}

func (value Value) String() string {
    return fmt.Sprint(value.value)
}

