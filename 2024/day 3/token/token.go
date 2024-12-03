package token

const (
	ILLEGAL = "ILLEGAL"
	INT     = "INT"
	EOF     = "EOF"
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
	"mul":   MUL,
	"do":    DO,
	"don't": DONT,
}

func LookupKeyword(word string) TokenType {
	if tok, ok := keywords[word]; ok {
		return tok
	}
	return ILLEGAL
}
