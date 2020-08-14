package golin2

import "testing"

func TestListToArray(t *testing.T) {
    var l = List{42, &List{1, nil}}
    var result = l.Array()
    
    if len(result) != 2 {
        t.Fatalf("Len of l.Array is %d expected 2", len(result));
    }
    
    if result[0] != 42 {
        t.Fatalf("l.Array()[0] is %d expected 42", result[0]);
    }
    
    if result[1] != 1 {
        t.Fatalf("l.Array()[1] is %d expected 42", result[1]);
    }
}

func TestEmptyListToArray(t *testing.T) {
    var l *List = nil
    var result = l.Array()
    
    if len(result) != 0 {
        t.Fatalf("Len of empty list is %d expected 0", len(result));
    }
}

func TestListFirstRest(t *testing.T) {
    var l = List{42, nil}
    
    if l.First() != 42 {
        t.Fatal("l.First() is not 42")
    }
    
    if l.Rest() != nil {
        t.Fatal("l.Rest() is not nil")
    }
}

func TestListWithTwoElements(t *testing.T) {    var l = List{42, &List{1, nil}}
    
    if l.First() != 42 {
        t.Fatal("l.First() is not 42")
    }
    
    if l.Rest() == nil {
        t.Fatal("l.Rest() is nil")
    }
    
    if l.Rest().First() != 1 {
        t.Fatal("l.Rest().First() is not 1")
    }
    
    if l.Rest().Rest() != nil {
        t.Fatal("l.Rest().Rest() is not null")
    }
}

func TestListFirstIsNil(t *testing.T) {
    var l = List{nil, nil}

    if l.First() != nil {
        t.Fatal("l.First() is not nil")
    }

    if l.Rest() != nil {
        t.Fatal("l.Rest() is not nil")
    }
}

func TestNilString(t *testing.T) {
    var l SExpr = nil

    if String(l) != "nil" {
        t.Fatal("l.String() is not nil")
    }
}

func TestSimpleListString(t *testing.T) {
    var l = List{42, nil}
    var result = l.String()

    if result != "(42)" {
        t.Fatalf("l.String() is %s, expeceted (42)", result)
    }
}

func TestListWithTwoElementsString(t *testing.T) {
    var l = List{42, &List{1, nil}}

    if l.String() != "(42 1)" {
        t.Fatal("l.String() is not (42 1)")
    }
}

func TestListWithNestedList(t *testing.T) {
    var l = List{List{42, nil}, &List{1, nil}}
    var result = l.String()

    if result != "((42) 1)" {
        t.Fatalf("l.String() is %s, expected ((42) 1)", result)
    }
}

var first Func = func (args []SExpr) SExpr {
    var l = args[0].(List)
    
    return l.First()
}

var cons Func = func (args []SExpr) SExpr {
    var head = args[0]
    var tail = args[1].(List)
    
    return List{head, &tail}
}

func TestEvaluateFirst(t *testing.T) {
    var f = Atom{"first"}
    var a = Atom{"a"}
    
    var env = NewEnv()
    
    f.Assoc(env, first)
    a.Assoc(env, List{42,nil})
    
    var l = List{f, &List{a, nil}}    
    var result = l.Evaluate(env)
    
    if result != 42 {
        t.Fatal("l.Evaluate(env) is not 42")
    }
}

func TestEvaluateCons(t *testing.T) {
    var f = Atom{"cons"}
    var a = Atom{"a"}
    var b = Atom{"b"}
    
    var env = NewEnv()
    
    f.Assoc(env, cons)
    a.Assoc(env, 1)
    b.Assoc(env, List{42,nil})
    
    var l = List{f, &List{a, &List{b, nil}}}    
    var result = l.Evaluate(env).(List)
    
    if result.String() != "(1 42)" {
        t.Fatalf("(cons a b) is %s expected (1 42)", result.String())
    }
}


