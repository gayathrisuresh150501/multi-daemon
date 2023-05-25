package daemon

import (
	// "errors"
	// "package/daemon"
	"testing"
)

// var ErrInvalidNodeCount = errors.New("invalid node count")

func TestRead(t *testing.T) {

	t.Run("ValidInput", func(t *testing.T) {
		expectedCount := 0
		expectedErr := error(nil)

		// Act
		count, err := Read()

		// Assert
		if count != expectedCount {
			t.Errorf("Expected count: %d, but got: %d", expectedCount, count)
		}
		if err != expectedErr {
			t.Errorf("Expected error: %v, but got: %v", expectedErr, err)
		}
	})

	t.Run("InvalidInput", func(t *testing.T) {
		expectedCount := 0
		expectedErr := ErrInvalidNodeCount

		// Act
		count, err := Read()

		// Assert
		if count != expectedCount {
			t.Errorf("Expected count: %d, but got: %d", expectedCount, count)
		}
		if err != expectedErr {
			t.Errorf("Expected error: %v, but got: %v", expectedErr, err)
		}
	})
	t.Run("NegativeInput", func(t *testing.T) {
		expectedCount := -3
		expectedErr := ErrInvalidNodeCount

		// Act
		count, err := Read()

		// Assert
		if count != expectedCount {
			t.Errorf("Expected count: %d, but got: %d", expectedCount, count)
		}
		if err != expectedErr {
			t.Errorf("Expected error: %v, but got: %v", expectedErr, err)
		}
	})

}
