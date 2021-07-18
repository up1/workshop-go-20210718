package db_test

import (
	"day01/db"
	"testing"
)

func TestSuccessCase(t *testing.T) {
	result, _ := db.SaveData("")
	if result != 100 {
		t.Errorf("Expect 100 but found %v", result)
	}
}

func BenchmarkRandInt(b *testing.B) {
    for i := 0; i < b.N; i++ {
        db.SaveData("")
    }
}