
package parser

type TokenType int

const (
    Name TokenType = iota
)

type Token interface {
    Value() string
    Type() TokenType
}

type Parser struct {
    text string
    length int
    position int
}

type token struct {
    value string
    typ TokenType
}

func (t token) Value() string { 
    return t.value
}

func (t token) Type() TokenType { 
    return t.typ
}

func NewParser(text string) Parser {
    return Parser{text, 0, 0}
}

func (parser Parser) NextToken() Token {
    if len(parser.text) == 0 {
        return nil
    }
    
    return token{parser.text, Name}
}

