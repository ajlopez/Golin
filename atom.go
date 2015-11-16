package golin

type Atom struct {
    name Symbol
}

func (a *Atom) Evaluate(env *Env) SExpr {
    return env.Find(a.name)
}

func (a *Atom) String() string {
    return string(a.name)
}



