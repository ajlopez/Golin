package golin

import "testing"

func TestEvaluateUndefinedAtom(t *testing.T) {
    var env = NewEnv()
    var atom = Atom{"foo"}
    
    var result = atom.Evaluate(env)
    
    if result != nil {
        t.Fatal("atom.Evaluate(env) is not nil")
    }
}

func TestAtomString(t *testing.T) {
    var atom = Atom{"foo"}
    
    var result = atom.String()
    
    if result != "foo" {
        t.Fatal("atom.String() is not foo")
    }
}
