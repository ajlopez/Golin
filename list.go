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

func asString(l *List) string {
    if l == nil {
        return ""
    }
    
    var shead = AsString(l.head)
    var stail = asString(l.tail)
    
    var result = " " + shead
    
    if stail != "" {
        result += " "
    }
    
    return result + stail
}

func (l *List) AsString() string {
    if l == nil {
        return "nil"
    }
    
    var result = "(" + AsString(l.head) + asString(l.tail) + ")"
    
    return result
}

