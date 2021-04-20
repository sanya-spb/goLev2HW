## Урок 4. Concurrency, часть 1: горутины и каналы. Пакет context

1. С помощью пула воркеров написать программу, которая запускает 1000 горутин, каждая из которых увеличивает число на 1. Дождаться завершения всех горутин и убедиться, что при каждом запуске программы итоговое число равно 1000. 
2. Написать программу, которая при получении в канал сигнала SIGTERM останавливается не позднее, чем за одну секунду (установить таймаут).

==== задание 1 
вариант 1 сделал не по ТЗ, но если есть что поругать, жду..

вариант 2 сделал универсальным

вариант 3 сделал по ТЗ


==== задание 2
вариант 1 сделано обычное завершение, без ожидания.


вариант 2 предусмотрено время на завершение. 
В зависимости от того в какой момент вызвать остановку - результат разный:
```go
$ go run srs_02_v2/main.go 
2021/04/20 16:42:38 doSomething() started
.........|^C.....2021/04/20 16:42:40 doSomething() stopped

$ go run srs_02_v2/main.go 
2021/04/20 16:42:46 doSomething() started
.........|.......^C..|2021/04/20 16:42:50 server has gracefully finished
2021/04/20 16:42:50 doSomething() stopped
```