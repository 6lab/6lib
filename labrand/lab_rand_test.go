package labrand

import (
	"testing"
)

func TestGetUUID(t *testing.T) {
	id := GetUUID()

	if len(id) != 36 {
		t.Error("Error with GetUUID()")
	}
}

func BenchmarkGetUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetUUID()
	}
}
