package main

import (
	"fmt"
	"time"
)

type ErrorWithTime struct {
	text string
	time time.Time
}

func New(text string) error {
	return &ErrorWithTime{
		text: text,
		time: time.Now(),
	}
}

func (e *ErrorWithTime) Error() string {
	return fmt.Sprintf("descr: %s\nat time: %s", e.text, e.time.String())
}

func main() {
	for i := 3; i >= -3; i-- {
		if result, err := divByMyself(i); err != nil {
			fmt.Printf("Error on: %[1]d / %[1]d\n%[2]s\n", i, err)
		} else {
			fmt.Printf("divByMyself: %[1]d / %[1]d = %[2]d\n", i, result)
		}
	}
}

// Делим "n" само на себя
func divByMyself(n int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			if n == 0 {
				err = New("На ноль делить не умею!")
			} else {
				err = New("Что-то пошло не так..")
			}
		}
	}()
	return n / n, nil
}
