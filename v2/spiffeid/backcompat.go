package spiffeid

func allowBackcompatTrustDomainChar(c uint8) bool {
	if isSubDelim(c) {
		return true
	}
	switch c {
	// unreserved
	case '~':
		return true
	default:
		return false
	}
}

func allowBackcompatPathChar(c uint8) bool {
	if isSubDelim(c) {
		return true
	}
	switch c {
	// unreserved
	case '~':
		return true
	// gen-delims
	case ':', '[', ']', '@':
		return true
	default:
		return false
	}
}

func isSubDelim(c uint8) bool {
	switch c {
	case '!', '$', '&', '\'', '(', ')', '*', '+', ',', ';', '=':
		return true
	default:
		return false
	}
}

func notAllowBackcompatTrustDomainChar(c uint8) bool {
	return false
}

func notAllowBackcompatPathChar(c uint8) bool {
	return false
}

func getTrustDomainCharFn(o *options) func(uint8) bool {
	isBackcompatTrustDomainCharFn := getDefaultTrustDomainCharFn()
	if o != nil && o.allowCharsetBackCompat {
		isBackcompatTrustDomainCharFn = allowBackcompatTrustDomainChar
	}
	return isBackcompatTrustDomainCharFn
}

func getIsPathCharFn(o *options) func(uint8) bool {
	isPathCharFn := getDefaultPathCharFn()
	if o != nil && o.allowCharsetBackCompat {
		isPathCharFn = allowBackcompatPathChar
	}
	return isPathCharFn
}
