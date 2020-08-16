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

func toString(l *List) string {
    if l == nil {
        return ""
    }
    
    var shead = l.head.String()
    var stail = toString(l.tail)
    
    var result = " " + shead
    
    if stail != "" {
        result += " "
    }
    
    return result + stail
}

func (l *List) Array() []SExpr {
    if l == nil {
        return make([]SExpr, 0)
    }
    
    if (l.tail == nil) {
        var arr = make([]SExpr, 1)
        arr[0] = l.head
        return arr
    }
    
    var arrtail = l.tail.Array()
    var arr = make([]SExpr, len(arrtail) + 1)
    arr[0] = l.head
    copy(arr[1:], arrtail)
    
    return arr
}

func (l List) String() string {
    var result = "(" + l.head.String() + toString(l.tail) + ")"
    
    return result
}

func (l List) Evaluate(env *Env) SExpr {
    var fn = l.head.Evaluate(env).(Func)
    
    var elems = l.tail.Array()
    var args []SExpr = make([]SExpr, len(elems))
    
    for k := 0; k < len(elems); k++ {
        args[k] = elems[k].Evaluate(env)
    }
    
    return fn(args)
}

