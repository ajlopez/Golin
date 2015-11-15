package golin

type SExpr interface {}

type SValue interface {
    Evaluate(*Env) SExpr
}

type Symbol string

func Evaluate(a SExpr, env *Env) SExpr {
    if v, ok := a.(SValue); ok {
        return v.Evaluate(env)
    }
    
    return a
}

