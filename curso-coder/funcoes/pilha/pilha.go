package main

import "runtime/debug"

func f3() {
	debug.PrintStack()
}

func f2() {
	f3()
}

func f1() {
	f2()
}

func main() {
	f1()

	/* Saída no console: Temos aqui uma pilha de chamadas de funções
		goroutine 1 [running]:
	runtime/debug.Stack(0xc000042778, 0xc000072f78, 0x4049a5)
		/usr/local/go/src/runtime/debug/stack.go:24 +0x9f
	runtime/debug.PrintStack()
		/usr/local/go/src/runtime/debug/stack.go:16 +0x25
	main.f3(...)
		/home/gabriel/programmation/study/GolangRepository/curso-coder/funcoes/pilha/pilha.go:6
	main.f2(...)
		/home/gabriel/programmation/study/GolangRepository/curso-coder/funcoes/pilha/pilha.go:10
	main.f1(...)
		/home/gabriel/programmation/study/GolangRepository/curso-coder/funcoes/pilha/pilha.go:14
	main.main()
		/home/gabriel/programmation/study/GolangRepository/curso-coder/funcoes/pilha/pilha.go:18 +0x25
	*/

}
