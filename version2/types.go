package golin2

type SExpr interface {
    Evaluate(*Env) SExpr
    String() string
}

type Func func (args []SExpr) SExpr

func (fn Func) Evaluate(env *Env) SExpr {
    return fn
}

func (fn Func) String() string {
    return "fn"
}