package main

import (
	"errors"
	"fmt"
)

func main() {
	for i := 3; i >= -3; i-- {
		if result, err := divByMyself(i); err != nil {
			fmt.Printf("%[2]s on: %[1]d / %[1]d\n", i, fmt.Errorf("%w", err))
		} else {
			fmt.Printf("divByMyself: %[1]d / %[1]d = %[2]d\n", i, result)
		}
	}
}

// Делим "n" само на себя
func divByMyself(n int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case error:
				err = x // тут наш случай, остальное пусть будет для полноты картины.
			case string:
				err = errors.New(x)
			default:
				err = errors.New("Unknown panic")
			}
		}
	}()
	return n / n, nil
}
