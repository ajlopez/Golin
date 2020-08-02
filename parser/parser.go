
package parser

type TokenType int

const (
    Name TokenType = iota
)

type Token interface {
    Value() string
    Type() TokenType
}

type Parser interface {
    NextToken() Token
}

type parser struct {
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

func (parser parser) NextToken() Token {
    if len(parser.text) == 0 {
        return nil
    }
    
    return token{parser.text, Name}
}

func NewParser(text string) Parser {
    return parser{text, 0, 0}
}

