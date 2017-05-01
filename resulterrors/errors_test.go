package resulterrors

import "testing"

func TestImmutableError_Type(t *testing.T) {
	var _ error = immutableError("type test")
}
