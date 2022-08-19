package greet

import (
	"errors"
	"fmt"
)

func Hello() {
	fmt.Println("Hello, starting...")
}

func Divide(x int, y int) (int, error) {
	if y > 0 {
		return x / y, nil
	}

	return 0, errors.New("divide by zero")
}
