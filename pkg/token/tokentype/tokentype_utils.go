package tokentype

func FromIdentifier(identifier string) TokenType {
	switch identifier {
	case "and":
		return AND
	case "class":
		return CLASS
	case "else":
		return ELSE
	case "false":
		return FALSE
	case "fun":
		return FUN
	case "for":
		return FOR
	case "if":
		return IF
	case "nil":
		return NIL
	case "or":
		return OR
	case "print":
		return PRINT
	case "return":
		return RETURN
	case "super":
		return SUPER
	case "this":
		return THIS
	case "true":
		return TRUE
	case "var":
		return VAR
	case "while":
		return WHILE
	default:
		return IDENTIFIER
	}
}

func ConstantLiteralValue(tokenType TokenType) interface{} {
	switch tokenType {
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return nil
	}
}
