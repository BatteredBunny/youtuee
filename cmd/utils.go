package cmd

func VerifyRune(char rune) bool {
	return (char >= 'a' && char <= 'z') ||
		(char >= 'A' && char <= 'Z') ||
		(char >= '0' && char <= '9') ||
		(char == '-') ||
		(char == '_')
}

const EXPECTED_ID_LENGTH = 11

func VerifyPath(id string) bool {
	if len(id) < EXPECTED_ID_LENGTH {
		return false
	}

	for i, char := range id {
		if i >= EXPECTED_ID_LENGTH {
			if char == '?' { // Only urls with ? will possibly be longer than 11 characters long
				// Verification complete
				break
			} else {
				return false
			}
		}

		if !VerifyRune(char) {
			return false
		}
	}

	return true
}
