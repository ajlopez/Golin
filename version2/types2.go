package golin2

type SExpr interface {
    Evaluate(*Env) *SExpr
    String() string
}

type Value interface {}

type Func func (args []*SExpr) *SExpr

type Symbol string

