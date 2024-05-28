package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Generator struct {
	out chan int
}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) Do(n int) <-chan int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	g.out = make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			g.out <- r.Intn(n)
		}

		close(g.out)
	}()

	return g.out
}

type RandGenerator struct {
}

func NewRandGenerator() *RandGenerator {
	return &RandGenerator{}
}

func (rg RandGenerator) Rand(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	return r.Intn(max-min+1) + min
}

func main() {
	//gen := NewGenerator()
	//
	//for num := range gen.Do(3) {
	//	fmt.Println(num)
	//}

	gen := NewRandGenerator()
	fmt.Println(gen.Rand(0, 10))
}
