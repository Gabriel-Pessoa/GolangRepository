// em go, por padrão, você passa uma cópia de valores para um método
// com ponteiro, você passa uma referência do objeto para que o método possa alterar essa referência
package main

import "fmt"

/* Revisão

- um ponteiro é representado por um "*"

*/

func inc1(n int) {
	n++
}

func inc2(n *int) {
	// *, nesse  caso, é usado para desreferenciar, ou seja,
	// ter acesso ao valor no qual o ponteiro aponta

	// o asterisco me dá acesso ao valor, e com o valor eu consigo incrementar a variável "n"
	*n++
}

func main() {
	n := 1

	// passando o parâmetro como valor, uma cópia do valor. Não irá causar alteração na variável 'n' da linha 24
	inc1(n)
	fmt.Println(n)

	// revisão: & usado para obter o endereço da variável
	inc2(&n)        // passando por referência e não por valor; o endereço variável 'n' da linha 24
	fmt.Println(&n) // referência ou endereço de memória
	fmt.Println(n)  // alteração no valor de n

	/* OBS:
	Não é indicado criar funções que alteram variáveis "globais", perde-se o controle das alterações, podendo gerar bugs e comporta-
	mentos inesperados
	*/
}
