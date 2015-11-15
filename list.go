package golin

type List struct {
    head SExpr
    tail *List
}

func (l *List) First() SExpr {
    return l.head
}

func (l *List) Rest() *List {
    return l.tail
}

func (l *List) AsString() string {
    if l == nil {
        return "nil"
    }
    
    var result = "(" + AsString(l.head) + ")"
    
    return result
}

