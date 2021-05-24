package main

import "fmt"

func main() {
	//Para definir esse slice(s) o go criou, internamente, um array que esse slice(s) irá referenciar. Todo slice referencia um array? Aparentemente, sim
	s := make([]int, 10) // definindo o tipo e o tamanho do slice
	s[9] = 12
	fmt.Println(s)

	// criando nova versão do slice(s)
	s = make([]int, 10, 20)        //o slice(s),agora, tem 10 posições -referencia 10 posições do array criado internamente-, e o array interno tem 20 posições
	fmt.Println(s, len(s), cap(s)) // função cap: pega a capacidade máxima do slice (tamanho do array interno por ele referenciado)

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0) // add mais elementos ao slice, cobrindo toda a capacidade do array interno

	fmt.Println(s, len(s), cap(s))

	// add mais 1 elemento acima do limite do array interno
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s)) // o comportamento padrão é o array dobrar de tamanho

}
