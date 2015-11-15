package golin

import "testing"

func TestListFirst(t *testing.T) {
    var l = List{42, nil}
    
    if l.First() != 42 {
        t.Fail()
    }
    
    if l.Rest() != nil {
        t.Fail()
    }
}