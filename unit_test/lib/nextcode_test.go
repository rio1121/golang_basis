package nextcode

import "testing"

func TestNextCode(t *testing.T) {
	result := NextCode("ABC")
	if result != "BCD" {
		t.Error("Expected BCD, got", result)
	}
}
