package golin

import "fmt"

type SExpr interface {}

type SValue interface {
    Evaluate(*Env) SExpr
    AsString() string
}

type Symbol string

func Evaluate(a SExpr, env *Env) SExpr {
    if v, ok := a.(SValue); ok {
        return v.Evaluate(env)
    }
    
    return a
}

func AsString(a SExpr) string {
    if v, ok := a.(SValue); ok {
        return v.AsString()
    }
    
    return fmt.Sprint(a)
}

