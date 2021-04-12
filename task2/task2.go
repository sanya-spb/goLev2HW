/*
Демонстрационный модуль

показывает применение документации
содержит одну тестовую функцию и один пример ее использования
*/
package task2

import (
	"errors"
)

// Делим "n" само на себя
func DivByMyself(n int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case error:
				err = x // тут наш случай, остальное для полноты картины.
			case string:
				err = errors.New(x)
			default:
				err = errors.New("Что-то пошло не так..")
			}
		}
	}()
	return n / n, nil
}
