package golin

type Func func (args []SExpr) SExpr

func (fn Func) Evaluate(env *Env) SExpr {
    return fn
}

func (fn Func) String() string {
    return "fn"
}