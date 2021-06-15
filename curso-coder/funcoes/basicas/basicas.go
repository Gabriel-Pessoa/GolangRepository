package main

import "fmt"

// func: palavra reservada para declarar uma função

/* estrutura básicas de funções:

func nomeDaFuncao(params) retornosEsperado {

}

*/

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s \n", p1, p2)
}

func f3() string {
	return "F3"
}

// parâmetros do mesmo tipo pode-se definir o tipo no último parâmetro
func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2) // função sprint não imprime no console, mas sim, retorna uma string (formatada, quebra de linha).
}

func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	f1()
	f2("Param1", "Param2")

	r3, r4 := f3(), f4("Param1", "Param2")
	fmt.Println(r3)
	fmt.Println(r4)

	r51, r52 := f5() // eu posso ignorar um valor com underline, mas não posso atribuir o valor dessa função apenas a uma variável

	fmt.Println("F5:", r51, r52)
}
