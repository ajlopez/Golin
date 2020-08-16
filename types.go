package golin

type SExpr interface {
    Evaluate(*Env) SExpr
    String() string
}

