// этот файл будем анализировать
package main

func foo() {
}

func foo1() {
	go foo()
}

func foo2() {
	go foo()
	go foo()
	go foo()
	go foo()
}
