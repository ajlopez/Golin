package golin

import "testing"

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
    var l *List = nil

    if l.String() != "nil" {
        t.Fatal("l.String() is not nil")
    }
}

func TestSimpleListString(t *testing.T) {
    var l = List{42, nil}

    if l.String() != "(42)" {
        t.Fatal("l.String() is not (42)")
    }
}

func TestListWithTwoElementsString(t *testing.T) {
    var l = List{42, &List{1, nil}}

    if l.String() != "(42 1)" {
        t.Fatal("l.String() is not (42 1)")
    }
}

