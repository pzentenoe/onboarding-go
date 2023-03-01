package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg = &sync.WaitGroup{}
	wg.Add(4)

	go imprimir(wg, "1")
	go imprimir(wg, "2")

	go imprimir(wg, "3")
	fmt.Println("4")

	go imprimir(wg, "5")
	wg.Wait()

}

func imprimir(wg *sync.WaitGroup, msg string) {
	fmt.Println(msg)
	wg.Done()
}
