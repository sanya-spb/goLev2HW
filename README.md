## Урок 5: Concurrency часть 2: основы типов данных из пакета sync

1. Напишите программу, которая запускает n потоков и дожидается завершения их всех 
2. Реализуйте функцию для разблокировки мьютекса с помощью defer
3. Протестируйте производительность операций чтения и записи на множестве действительных чисел, безопасность которого обеспечивается sync.Mutex и sync.RWMutex для разных вариантов использования: 10% запись, 90% чтение; 50% запись, 50% чтение; 90% запись, 10% чтение

#### задание 1
все решение в коде, для ~~забавы~~ информативности добавил поток с анимашкой

#### задание 2
взял пример из методички, и подогнал его под решение задания

#### задание 3
```bash
$ go test -bench=. ./srs_03/main_test.go 
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-3570K CPU @ 3.40GHz
BenchmarkRunRWMutexMap_01-4         2828            445935 ns/op
BenchmarkRunRWMutexMap_05-4         1915            614964 ns/op
BenchmarkRunRWMutexMap_09-4         1605            627308 ns/op
BenchmarkRunMutexMap_01-4           2842            441621 ns/op
BenchmarkRunMutexMap_05-4           2101            537223 ns/op
BenchmarkRunMutexMap_09-4           1725            618434 ns/op
PASS
ok      command-line-arguments  7.261s
```