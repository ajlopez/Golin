package golin2

type SExpr interface {
    Evaluate(*Env) SExpr
    String() string
}

type Func func (args []SExpr) SExpr

