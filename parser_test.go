package golin

import "testing"

func TestParserReturnsNilIfEmptyString(t *testing.T) {
    parser := NewParser("")
    
    result := parser.NextToken()
    
    if result != nil {
        t.Fatal("Parser does not return nil on empty string")
    }
}

