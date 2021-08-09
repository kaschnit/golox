package stringutil

func IsRuneAlpha(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_'
}

func IsRuneNumeric(r rune) bool {
	return r >= '0' && r <= '9'
}

func IsRuneAlphaNumeric(r rune) bool {
	return IsRuneAlpha(r) || IsRuneNumeric(r)
}
