package token

type TokenType int

const (
	EOF TokenType = iota
	EOL
	Keyword
	QuotedString
	Variable
	BlockStart
	BlockEnd
	Semicolon
	Comment
	Illegal
	Regex
)

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

func (t Token) Lit(literal string) Token {
	t.Literal = literal
	return t
}

func (t Token) Is(typ TokenType) bool {
	return t.Type == typ
}

func (t Token) IsParameterEligible() bool {
	return t.Is(Keyword) || t.Is(QuotedString) || t.Is(Variable) || t.Is(Regex)
}
