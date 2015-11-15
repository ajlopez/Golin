package golin

type Env struct {
    values map[Symbol]SExpr
}

func (e *Env) Find(symbol Symbol) SExpr {
    return e.values[symbol]
}

