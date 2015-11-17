package golin

import "fmt"

type SExpr interface {}

type SValue interface {
    Evaluate(*Env) SExpr
    String() string
}

type Func func (args []SExpr) SExpr

type Symbol string

func Evaluate(a SExpr, env *Env) SExpr {
    if v, ok := a.(SValue); ok {
        return v.Evaluate(env)
    }
    
    return a
}

func String(a SExpr) string {
    if v, ok := a.(SValue); ok {
        return v.String()
    }
    
    return fmt.Sprint(a)
}

