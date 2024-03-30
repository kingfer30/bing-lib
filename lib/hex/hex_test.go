package hex_test

import (
	"testing"

	"github.com/kingfer30/bing-lib/lib/hex"
)

func TestHex(t *testing.T) {
	t.Log(hex.NewHex(32))
	t.Log(hex.NewHexLowercase(32))
}
