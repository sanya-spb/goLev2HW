## Урок 2: Кросс-компиляция в Go и более глубокое погружение в тулчейн

1. Выполните сборку ваших предыдущих программ под операционную систему, отличающуюся от текущей. Проанализируйте вывод команды file для полученного исполняемого файла. Попробуйте запустить исполняемый файл
2. Напишите документацию для одной из предыдущих программ. Запустите сервер документации локально. Убедитесь, что документация отображается корректно.

```bash
$ make help
 Choose a command run in goLev2HW:
  hw      task-02 HomeWork
  win64   Compile for win64 platform
  win32   Compile for win32 platform
  doc     Documentation
  clean   Clean build files.
  help    Show this
$ make hw
go run main.go
divByMyself: 3 / 3 = 1
divByMyself: 2 / 2 = 1
divByMyself: 1 / 1 = 1
Error on: 0 / 0
runtime error: integer divide by zero
divByMyself: -1 / -1 = 1
divByMyself: -2 / -2 = 1
divByMyself: -3 / -3 = 1
$ make win32
GOOS=windows GOARCH=386 go build -o app_main32.exe main.go
file app_main32.exe
app_main32.exe: PE32 executable (console) Intel 80386 (stripped to external PDB), for MS Windows
wine app_main32.exe
divByMyself: 3 / 3 = 1
divByMyself: 2 / 2 = 1
divByMyself: 1 / 1 = 1
Error on: 0 / 0
runtime error: integer divide by zero
divByMyself: -1 / -1 = 1
divByMyself: -2 / -2 = 1
divByMyself: -3 / -3 = 1
$ make win64
GOOS=windows GOARCH=amd64 go build -o app_main64.exe main.go
file app_main64.exe
app_main64.exe: PE32+ executable (console) x86-64 (stripped to external PDB), for MS Windows
wine app_main64.exe
divByMyself: 3 / 3 = 1
divByMyself: 2 / 2 = 1
divByMyself: 1 / 1 = 1
Error on: 0 / 0
runtime error: integer divide by zero
divByMyself: -1 / -1 = 1
divByMyself: -2 / -2 = 1
divByMyself: -3 / -3 = 1
$ make doc
go doc github.com/sanya-spb/goLev2HW/task2
package task2 // import "github.com/sanya-spb/goLev2HW/task2"

Демонстрационный модуль

показывает применение документации содержит одну тестовую функцию и один
пример ее использования

func DivByMyself(n int) (result int, err error)
go doc github.com/sanya-spb/goLev2HW/task2.DivByMyself
package task2 // import "github.com/sanya-spb/goLev2HW/task2"

func DivByMyself(n int) (result int, err error)
    Делим "n" само на себя
```