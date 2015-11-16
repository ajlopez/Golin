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

func TestAssocAndEvaluateAtom(t *testing.T) {
    var env = NewEnv()
    var atom = Atom{"foo"}
    atom.Assoc(env, 42);
    
    var result = atom.Evaluate(env)
    
    if result != 42 {
        t.Fatal("atom.Evaluate(env) is not 42")
    }
}

func TestAtomString(t *testing.T) {
    var atom = Atom{"foo"}
    
    var result = atom.String()
    
    if result != "foo" {
        t.Fatal("atom.String() is not foo")
    }
}
