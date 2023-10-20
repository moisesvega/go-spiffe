//go:build !spiffeid_charset_backcompat
// +build !spiffeid_charset_backcompat

package spiffeid

func getDefaultTrustDomainCharFn() func(c uint8) bool {
	return notAllowBackcompatTrustDomainChar
}

func getDefaultPathCharFn() func(c uint8) bool {
	return notAllowBackcompatPathChar
}
