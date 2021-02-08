package lib

import (
	"example.com/owner/repo/testutil"
	"testing"
)

func TestFoo(t *testing.T) {
	_ = GetFoo()
	_ = testutil.GetFakeFoo()
}
