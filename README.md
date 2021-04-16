## Урок 3: Работа с внешними зависимостями

Для закрепления усвоенного материала повторите шаги раздела «Практическая часть» самостоятельно для своего проекта:
Создайте новый проект с использованием инструментария go mod.
1. Опубликуйте проект в репозитории, установив номер версии, указывающий на активный этап разработки библиотеки.
2. Обновите номера версий зависимостей в библиотеке.
3. Сделайте изменения в проекте и запушьте их с мажорным обновлением версии пакета.
4. Очистите неиспользуемые библиотеки.


запуск без тагов semver
```bash
$ ./app_main 
version: {Version:devel Commit:git-3a3d567 BuildTime:2021-04-16_07:02:55 Copyright:sanya-spb}
divByMyself: 1 / 1 = 1
Error on: 0 / 0
runtime error: integer divide by zero
divByMyself: -1 / -1 = 1
```

далее выполним пункт 1
```bash
$ git tag -a v0.0.1 -m 'pre-release.'
$ make build
$ ./app_main 
version: {Version:v0.0.1 Commit:git-3a3d567 BuildTime:2021-04-16_07:20:43 Copyright:sanya-spb}
divByMyself: 1 / 1 = 1
Error on: 0 / 0
runtime error: integer divide by zero
divByMyself: -1 / -1 = 1
```