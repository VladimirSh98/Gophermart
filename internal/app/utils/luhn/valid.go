package luhn

import "unicode"

func IsValid(s string) bool {
	var sum int
	var alt bool

	for i := len(s) - 1; i >= 0; i-- {
		r := rune(s[i])
		if !unicode.IsDigit(r) {
			return false
		}
		n := int(r - '0')
		if alt {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
		alt = !alt
	}

	return sum%10 == 0
}
