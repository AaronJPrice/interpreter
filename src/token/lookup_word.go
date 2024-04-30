package token

// Map token literal to TokenType
func LookupWord(word string) TokenType {
	switch word {
	case "fn":
		return FUNCTION
	case "let":
		return LET
	case "true":
		return TRUE
	case "false":
		return FALSE
	case "if":
		return IF
	case "else":
		return ELSE
	case "return":
		return RETURN
	default:
		return IDENT
	}
}
