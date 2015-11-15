package golin

type List struct {
    head SExpr
    tail SExpr
}

func (l *List) First() SExpr {
    return l.head
}

func (l *List) Rest() SExpr {
    return l.tail
}

