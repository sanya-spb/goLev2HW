PROJECT?=github.com/sanya-spb/goLev2HW
PROJECTNAME=$(shell basename "$(PROJECT)")

## hw: task-02 HomeWork
hw:
	go run main.go

## win64: Compile for win64 platform
win64:
	GOOS=windows GOARCH=amd64 go build -o app_main64.exe main.go
	file app_main64.exe
	wine app_main64.exe

## win32: Compile for win32 platform
win32:
	GOOS=windows GOARCH=386 go build -o app_main32.exe main.go
	file app_main32.exe
	wine app_main32.exe

## doc: Documentation
doc:
	go doc github.com/sanya-spb/goLev2HW/task2
	go doc github.com/sanya-spb/goLev2HW/task2.DivByMyself

## clean: Clean build files.
clean: 
	go clean
	rm app_main*

## help: Show this
help: Makefile
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'