package main

import (
	"fmt"
	"testing"
)

func Test_Ok(t *testing.T) {
	hash := do(1, 2, 3, 4.1, "qwe", false, 1+1i)
	if "446ed5629dd0ed769009bec6ace94192807d226d2ea66fdb3e5f221edd5a7266" != fmt.Sprintf("%x", hash) {
		t.Error()
	}
}
