package integers

import "testing"

func TestAdders(t *testing.T) {
	t.Run("should add and return 4", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("expected '%d' bu got '%d'", expected, sum)
		}
	})
}
