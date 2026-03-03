package internal

func verifyRune(char rune) bool {
	return (char >= 'a' && char <= 'z') ||
		(char >= 'A' && char <= 'Z') ||
		(char >= '0' && char <= '9') ||
		(char == '-') ||
		(char == '_')
}

const expected_youtube_video_id_length = 11

func verifyPath(id string) bool {
	if len(id) < expected_youtube_video_id_length {
		return false
	}

	for i, char := range id {
		if i >= expected_youtube_video_id_length {
			if char == '?' { // Only urls with ? will possibly be longer than 11 characters long
				// Verification complete
				break
			} else {
				return false
			}
		}

		if !verifyRune(char) {
			return false
		}
	}

	return true
}
