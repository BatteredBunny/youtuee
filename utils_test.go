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
