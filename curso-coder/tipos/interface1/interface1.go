package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// interfaces são implementadas implicitamente, não utilizam a palavra reservada implements.
// O compilador infere que a struct implementa a interface, se esta struct tiver um "método"
// descrito na interface

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

// recebe como parâmetro uma "classe" que respeita, ou melhor, que implementa a interface imprimivel
func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	// polimorfismo: qualidade de assumir diferentes formas.
	// Em O.O.: referências de tipos de classes mais abstratas representem o comportamento das classes concretas que referenciam
	/* Ex:
				public abstract class OperacaoMatematica {
		    		public abstract double calcular(double x, double y);
				}

				public class Soma extends OperacaoMatematica {
	  				public double calcular(double x, double y) {
	        			return x + y;
	    			}
				}

				public class Subtracao extends OperacaoMatematica {
					public double calcular(double x, double y) {
						return x - y;
					}
				}
	*/

	var coisa imprimivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Calça Jeans", 79.90}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	p2 := produto{"Camisa", 59.90}
	imprimir(p2)
}
