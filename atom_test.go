package golin

import "testing"

func TestEvaluateUndefinedAtom(t *testing.T) {
    var env = NewEnv()
    var atom = Atom{"foo"}
    
    var result = atom.Evaluate(env)
    
    if result != nil {
        t.Fail()
    }
}
