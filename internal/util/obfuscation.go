package util

func ValidateDomainObfuscation(domain, obfuscationDomain string) bool {
	if len(domain) != len(obfuscationDomain) {
		return false
	}

	for i := 0; i < len(domain); i++ {
		dRune := []rune(domain)[i]
		odRune := []rune(obfuscationDomain)[i]

		switch dRune {
		case '.':
			if odRune != '.' {
				return false
			}
		default:
			switch odRune {
			case dRune:
				continue
			case '*':
				continue
			default:
				return false
			}
		}
	}

	return true
}
