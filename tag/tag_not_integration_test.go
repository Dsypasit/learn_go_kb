//go:build !integration

package tag

import "testing"

func TestTag4(t *testing.T) {
	t.Log("with build tag not integration")
}
