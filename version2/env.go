package golin2

type Env struct {
    values map[string]SExpr
}

func NewEnv() *Env {
    var e = new(Env)
    e.values = make(map[string]SExpr)
    
    return e
}

func (e *Env) Set(name string, value SExpr) {
    e.values[name] = value
}

func (e *Env) Get(name string) SExpr {
    return e.values[name]
}

