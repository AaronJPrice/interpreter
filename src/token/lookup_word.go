package token

func LookupWord(word string) string {
	switch word {
	case "fn":
		return FUNCTION
	case "let":
		return LET
	default:
		return IDENT
	}
}
