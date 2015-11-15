package golin

import "testing"

func TestListFirst(t *testing.T) {
    var l = List{42, nil}
    
    if l.First() != 42 {
        t.Fatal("l.First() is not 42")
    }
    
    if l.Rest() != nil {
        t.Fatal("l.Rest() is not nil")
    }
}