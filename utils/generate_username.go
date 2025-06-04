package utils

import (
	"fmt"
	"math/rand"
)

func GenerateUsername() string {
	var (
		ADJECTIVES = []string{
			"silent", "swift", "focused", "precise", "stealthy", "clever", "debugging",
			"minimal", "atomic", "zen", "furious", "clean", "crisp", "agile", "refactored",
			"sharp", "restless", "tactical", "curious", "lazy", "async", "pure", "stateless",
			"encrypted", "modular", "pragmatic", "wired", "recursive", "snappy", "lean",
			"faultless", "isolated", "cleanroom", "composed", "terminal", "invisible", "headless",
		}

		NOUNS = []string{
			"ninja", "dev", "coder",
		}
	)

	adj := ADJECTIVES[rand.Intn(len(ADJECTIVES))]
	noun := NOUNS[rand.Intn(len(NOUNS))]
	return fmt.Sprintf("%s_%s", adj, noun)
}
