package internal

import (
	"testing"

	assert "github.com/BatteredBunny/testingassert"
)

func TestCharSet(t *testing.T) {
	assert.TestState = t
	assert.HideSuccess = true

	// Alphanumeric characters should get through
	assert.Assert(verifyRune('a'), "Normal character fails")
	assert.Assert(verifyRune('z'), "Normal character fails")
	assert.Assert(verifyRune('A'), "Normal character fails")
	assert.Assert(verifyRune('Z'), "Normal character fails")
	assert.Assert(verifyRune('0'), "Number fails")
	assert.Assert(verifyRune('9'), "Number fails")

	// Unwanted special characters
	assert.Assert(!verifyRune('['), "Unwanted special character gets through")
	assert.Assert(!verifyRune('.'), "Unwanted special character gets through")
	assert.Assert(!verifyRune('!'), "Unwanted special character gets through")

	// Wanted special character
	assert.Assert(verifyRune('_'), "Wanted special char fails")
	assert.Assert(verifyRune('-'), "Wanted special char fails")
}

func TestIDVerification(t *testing.T) {
	assert.TestState = t
	assert.HideSuccess = true

	assert.Assert(verifyPath("gocwRvLhDf8"), "Normal bare video fails to pass verification")
	assert.Assert(verifyPath("uQx8jKiIDTI?si=cF4RjcTGM1cGx7JM"), "Video ID with referral info fails to pass verification")

	assert.Assert(!verifyPath("goacwRvLhDf8"), "Too long malformed ID gets through")
	assert.Assert(!verifyPath("hi"), "Too short id")
	assert.Assert(!verifyPath("favicon.ico"), "Bot junk traffic gets through")
	assert.Assert(!verifyPath("page.php"), "Bot junk traffic gets through")
}
