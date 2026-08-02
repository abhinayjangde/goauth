package main

import (
	"errors"
	"fmt"
	"log/slog"
)

func divide(a, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}
func main() {

	result, err := divide(4, 0)

	if err != nil {
		slog.Error(err.Error())
	}
	fmt.Println("result is", result)
}
