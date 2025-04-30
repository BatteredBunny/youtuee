package main

import (
	"testing"

	assert "github.com/BatteredBunny/testingassert"
)

func TestCharSet(t *testing.T) {
	assert.TestState = t
	assert.HideSuccess = true

	// Alphanumeric characters should get through
	assert.Assert(VerifyRune('a'), "Normal character fails")
	assert.Assert(VerifyRune('z'), "Normal character fails")
	assert.Assert(VerifyRune('A'), "Normal character fails")
	assert.Assert(VerifyRune('Z'), "Normal character fails")
	assert.Assert(VerifyRune('0'), "Number fails")
	assert.Assert(VerifyRune('9'), "Number fails")

	// Unwanted special characters
	assert.Assert(!VerifyRune('['), "Unwanted special character gets through")
	assert.Assert(!VerifyRune('.'), "Unwanted special character gets through")
	assert.Assert(!VerifyRune('!'), "Unwanted special character gets through")

	// Wanted special character
	assert.Assert(VerifyRune('_'), "Wanted special char fails")
	assert.Assert(VerifyRune('-'), "Wanted special char fails")
}

func TestIDVerification(t *testing.T) {
	assert.TestState = t
	assert.HideSuccess = true

	assert.Assert(VerifyPath("gocwRvLhDf8"), "Normal bare video fails to pass verification")
	assert.Assert(VerifyPath("uQx8jKiIDTI?si=cF4RjcTGM1cGx7JM"), "Video ID with referral info fails to pass verification")

	assert.Assert(!VerifyPath("goacwRvLhDf8"), "Too long malformed ID gets through")
	assert.Assert(!VerifyPath("hi"), "Too short id")
	assert.Assert(!VerifyPath("favicon.ico"), "Bot junk traffic gets through")
	assert.Assert(!VerifyPath("page.php"), "Bot junk traffic gets through")
}
