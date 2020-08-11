package golin2

type Atom struct {
    name Symbol
}

func (a *Atom) Evaluate(env *Env) *SExpr {
    return env.Get(a.name)
}

func (a *Atom) Assoc(env *Env, value SExpr) {
    env.Set(a.name, &value)
}

func (a *Atom) String() string {
    return string(a.name)
}

