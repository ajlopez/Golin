package golin2

type Env struct {
    values map[Symbol]*SExpr
}

func NewEnv() *Env {
    var e = new(Env)
    e.values = make(map[Symbol]*SExpr)
    
    return e
}

func (e *Env) Set(symbol Symbol, value *SExpr) {
    e.values[symbol] = value
}

func (e *Env) Get(symbol Symbol) *SExpr {
    return e.values[symbol]
}

