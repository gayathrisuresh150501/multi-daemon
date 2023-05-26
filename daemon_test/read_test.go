package daemon_test

import (
	// "errors"
	"testing"
	"package/daemon"
)

// var ErrInvalidNodeCount = errors.New("invalid node count")

func TestRead(t *testing.T) {

	t.Run("Valid Input", func(t *testing.T) {
		expectedCount := 2
		expectedErr := error(nil)

		count, err := daemon.Read()

		if count != expectedCount {
			t.Errorf("Expected count: %d, but got: %d", expectedCount, count)
		}
		if err != expectedErr {
			t.Errorf("Expected error: %v, but got: %v", expectedErr, err)
		}
	})

	t.Run("Invalid Input", func(t *testing.T) {
		expectedCount := 0
		expectedErr := daemon.ErrInvalidNodeCount

		count, err := daemon.Read()

		if count != expectedCount {
			t.Errorf("Expected count: %d, but got: %d", expectedCount, count)
		}
		if err != expectedErr {
			t.Errorf("Expected error: %v, but got: %v", expectedErr, err)
		}
	})
	t.Run("Negative Input", func(t *testing.T) {
		expectedCount := -3
		expectedErr := daemon.ErrInvalidNodeCount

		count, err := daemon.Read()

		if count != expectedCount {
			t.Errorf("Expected count: %d, but got: %d", expectedCount, count)
		}
		if err != expectedErr {
			t.Errorf("Expected error: %v, but got: %v", expectedErr, err)
		}
	})

}
