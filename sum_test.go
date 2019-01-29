package sum

import "testing"

func TestSum(t *testing.T) {
	total := Add(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestSumLarge(t *testing.T) {
	total := Add(500, 500)
	if total != 1000 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 1000)
	}
}
