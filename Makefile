PROJECT?=github.com/sanya-spb/goLev2HW
PROJECTNAME=$(shell basename "$(PROJECT)")

## task1
task1:
	@echo "_________________________________________"
	@echo "_______________ task 01   _______________"
	go run ./srs_01/main.go

## task2
task2:
	@echo "_________________________________________"
	@echo "_______________ task 02   _______________"
	@echo "_________________________________________"
	@echo "_______________ variant a _______________"
	go run ./srs_02a/main.go
	@echo "_________________________________________"
	@echo "_______________ variant b _______________"
	go run ./srs_02b/main.go

## task3
task3:
	@echo "_________________________________________"
	@echo "_______________ task 03   _______________"
	go run ./srs_03/main.go

## task4
task4:
	@echo "_________________________________________"
	@echo "_______________ task 04   _______________"
	go run ./srs_04/main.go

## clean: Clean build files.
clean: 
	go clean
	find . -name "out_*" -print0 | xargs -0 rm

## help: Show this
help: Makefile
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'