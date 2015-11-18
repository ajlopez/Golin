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
    
    var shead = String(l.head)
    var stail = toString(l.tail)
    
    var result = " " + shead
    
    if stail != "" {
        result += " "
    }
    
    return result + stail
}

func (l List) String() string {
    var result = "(" + String(l.head) + toString(l.tail) + ")"
    
    return result
}

func (l List) Evaluate(env *Env) SExpr {
    var fn = Evaluate(l.head, env).(Func)
    var args []SExpr
    var arg = Evaluate(l.tail.First(), env)
    args = append(args, arg)
    
    return fn(args)
}

