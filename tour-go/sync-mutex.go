// se nós apenas quiséssemos ter certeza que uma única goroutine pode acessar uma variável de cada vez para evitar conflitos
package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	//Podemos definir um bloco de código a ser executado em exclusão mútua pelo que o rodeia com uma chamada para Lock e Unlock como mostrado abaixo
	c.mu.Lock()
	// Bloqueie para que apenas uma goroutine por vez possa acessar o mapa c.v.
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock() //podemos usar defer para garantir que o mutex será desbloqueado no método

	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 100; i++ {
		go c.Inc("algumakey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("algumakey"))
}
