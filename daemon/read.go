package daemon

// import "fmt"

func Read() (int, error) {
	var instCount int = 2
		// fmt.Scanf("%d", &instCount)

	if instCount <= 0 {
		return instCount, ErrInvalidNodeCount
	}

	return instCount, nil
}
