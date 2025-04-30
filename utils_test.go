package main

import (
	"testing"

	assert "github.com/BatteredBunny/testingassert"
)

func TestFormatDescription(t *testing.T) {
	assert.TestState = t
	assert.HideSuccess = true
	assert.Equals(FormatDescription("🌎 Get an exclusive 15% discount on Saily data plans! Use code bigmoney at checkout. Download Saily app or go to https://saily.com/bigmoney\n\nhttp://www.neongrizzly.com to buy Erik's shirt\n\n\"Big Money Salvia Theme\" by Hot Dad http://www.youtube.com/HotDad\n\nGentleman Erik Intro Pic by http://www.twitter.com/SentientPizzaB\n\nhttp://www.patreon.com/commentiquette\nhttp://www.twitter.com/commentiquette"), "🌎 Get an exclusive 15% discount on Saily data plans! Use code bigmoney at checkout. Download Saily app or go to https://saily.com/bigmoneyhttp://www.neongri...")
}

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
