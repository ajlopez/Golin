
package golin

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

func NewParser(text string) Parser {
    return Parser{text, 0, 0}
}

func (parser Parser) NextToken() Token {
    return nil
}

