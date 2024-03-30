package hex_test

import (
	"testing"

	"github.com/kingfer30/bing-lib/lib/hex"
)

func TestUUID(t *testing.T) {
	t.Log(hex.NewUUID())
}
