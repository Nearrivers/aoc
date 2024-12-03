package token

const (
	ILLEGAL = "ILLEGAL"
	INT     = "INT"
	CRLF    = "CRLF"
	LPARENT = "("
	RPARENT = ")"
	COMMA   = ","
	MUL     = "mul"
	DO      = "do"
	DONT    = "dont"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"mul": MUL,
}

func LookupKeyword(word string) TokenType {
	if tok, ok := keywords[word]; ok {
		return tok
	}
	return ILLEGAL
}
