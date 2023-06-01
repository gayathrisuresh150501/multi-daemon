package daemon

// import "fmt"
var InstCount int = 2
func Read() (int, error) {
		// fmt.Scanf("%d", &instCount)

	if InstCount <= 0 {
		return InstCount, ErrInvalidNodeCount
	}

	return InstCount, nil
}
