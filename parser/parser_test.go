package parser

import "testing"

func TestParserReturnsNilIfEmptyString(t *testing.T) {
    parser := NewParser("")
    
    result := parser.NextToken()
    
    if result != nil {
        t.Fatal("Parser does not return nil on empty string")
    }
}

func TestParseName(t *testing.T) {
    parser := NewParser("name")
    
    result := parser.NextToken()
    
    if result == nil || result.Value() != "name" || result.Type() != Name {
        t.Fatal("Parser does not parse name")
    }
}
