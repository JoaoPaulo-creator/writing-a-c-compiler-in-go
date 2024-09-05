package token

type TokenType string

const (
	OPEN_PAREN  = "("
	CLOSE_PAREN = ")"
	OPEN_BRACE  = "{"
	CLOSE_BRACE = "}"

	SEMICOLON = ";"
	INT       = "int"
	MAIN      = "main"
	VOID      = "void"
	RETURN    = "return"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"int":    INT,
	"void":   VOID,
	"main":   MAIN,
	"return": RETURN,
}
