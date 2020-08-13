package golin2

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
    var value = Atom{"bar"}
    var _ SExpr = (*Atom)(nil)
    atom.Assoc(env, value);
    
    var result SExpr = atom.Evaluate(env)
    
    if result.String() != "bar" {
        t.Fatal("atom.Evaluate(env) is not value")
    }
}

func TestAtomString(t *testing.T) {
    var atom = Atom{"foo"}
    
    var result = atom.String()
    
    if result != "foo" {
        t.Fatal("atom.String() is not foo")
    }
}

