package filter

import (
	"strings"
	"testing"

	"github.com/DAO-Metaplayer/aiops-bundler/internal/testutils"
)

func TestIsValidAiOpHash(t *testing.T) {
	if ok := IsValidAiOpHash(testutils.MockHash); !ok {
		t.Fatalf("%s: got false, want true", testutils.MockHash)
	}

	allNumHash := strings.ReplaceAll(testutils.MockHash, "dead", "0101")
	if ok := IsValidAiOpHash(allNumHash); !ok {
		t.Fatalf("%s: got false, want true", allNumHash)
	}
}

func TestIsValidAiOpHashAllCaps(t *testing.T) {
	hash := strings.ToUpper(testutils.MockHash)
	if ok := IsValidAiOpHash(hash); !ok {
		t.Fatalf("%s: got false, want true", hash)
	}
}

func TestIsValidAiOpHashEmptyString(t *testing.T) {
	hash := ""
	if ok := IsValidAiOpHash(hash); ok {
		t.Fatalf("%s: got true, want false", hash)
	}
}

func TestIsValidAiOpHashEmptyHexString(t *testing.T) {
	hash := "0x"
	if ok := IsValidAiOpHash(hash); ok {
		t.Fatalf("%s: got true, want false", hash)
	}
}

func TestIsValidAiOpHashNoPrefix(t *testing.T) {
	hash := strings.TrimPrefix(testutils.MockHash, "0x")
	if ok := IsValidAiOpHash(hash); ok {
		t.Fatalf("%s: got true, want false", hash)
	}
}

func TestIsValidAiOpHashTooShort(t *testing.T) {
	hash := "0xdead"
	if ok := IsValidAiOpHash(hash); ok {
		t.Fatalf("%s: got true, want false", hash)
	}
}

func TestIsValidAiOpHashTooLong(t *testing.T) {
	hash := testutils.MockHash + "dead"
	if ok := IsValidAiOpHash(hash); ok {
		t.Fatalf("%s: got true, want false", hash)
	}
}

func TestIsValidAiOpHashInvalidChar(t *testing.T) {
	hash := strings.ReplaceAll(testutils.MockHash, "dead", "zzzz")
	if ok := IsValidAiOpHash(hash); ok {
		t.Fatalf("%s: got true, want false", hash)
	}
}
