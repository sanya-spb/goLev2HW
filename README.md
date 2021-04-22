## Урок 5: Concurrency часть 2: основы типов данных из пакета sync

1. Напишите программу, которая запускает n потоков и дожидается завершения их всех 
2. Реализуйте функцию для разблокировки мьютекса с помощью defer
3. Протестируйте производительность операций чтения и записи на множестве действительных чисел, безопасность которого обеспечивается sync.Mutex и sync.RWMutex для разных вариантов использования: 10% запись, 90% чтение; 50% запись, 50% чтение; 90% запись, 10% чтение

#### задание 1
все решение в коде, для ~~забавы~~ информативности добавил поток с анимашкой

#### задание 2
взял пример из методички, и подогнал его под решение задания

#### задание 3

вариант 1 - это рефакторинг кода преподавателя
```bash
$ go test -bench=. ./srs_03_v1/main_test.go 
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-3570K CPU @ 3.40GHz
BenchmarkRunRWMutexMap_01-4         2493            444102 ns/op
BenchmarkRunRWMutexMap_05-4         2060            551420 ns/op
BenchmarkRunRWMutexMap_09-4         1924            621033 ns/op
BenchmarkRunMutexMap_01-4           2268            440935 ns/op
BenchmarkRunMutexMap_05-4           1942            517074 ns/op
BenchmarkRunMutexMap_09-4           1970            609245 ns/op
PASS
ok      command-line-arguments  8.914s
```

вариант 2 - это моя "отсебятина"
```bash
$ go test -bench=. ./srs_03_v2/NewMapRWMu_test.go 
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-3570K CPU @ 3.40GHz
BenchmarkRunRWMutexMap_01-4         2503            440053 ns/op
BenchmarkRunRWMutexMap_05-4         2133            562686 ns/op
BenchmarkRunRWMutexMap_09-4         1890            643117 ns/op
PASS
ok      command-line-arguments  4.648s

$ go test -bench=. ./srs_03_v2/NewMapMu_test.go 
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-3570K CPU @ 3.40GHz
BenchmarkRunMutexMap_01-4           2420            479307 ns/op
BenchmarkRunMutexMap_05-4           2221            477369 ns/op
BenchmarkRunMutexMap_09-4           2392            498821 ns/op
PASS
ok      command-line-arguments  4.445s
```

> Вывод: оба варианта показывают одинаковую зависимость.  
> RWMutex становится "тяжелее" при более частых записях, а Mutex не зависит от преобладания записи над чтением