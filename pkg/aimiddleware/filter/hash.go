package filter

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	aiOpHashStrLen = 64
	aiOpHashRegex  = regexp.MustCompile(fmt.Sprintf("(?i)0x[0-9a-f]{%d}", aiOpHashStrLen))
)

func IsValidAiOpHash(aiOpHash string) bool {
	return len(strings.TrimPrefix(strings.ToLower(aiOpHash), "0x")) == aiOpHashStrLen &&
		aiOpHashRegex.MatchString(aiOpHash)
}
