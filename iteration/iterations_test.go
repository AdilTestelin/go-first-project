package iteration

import "testing"

func TestIterations(t *testing.T) {
	t.Run("should repart 'a' 5 times", func(t *testing.T) {
		expected := "aaaaa"
		actual := RepeatChar("a", 5)

		if actual != expected {
			t.Errorf("Expected '%q' but got '%q'", expected, actual)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatChar("a", 5)
	}
}
