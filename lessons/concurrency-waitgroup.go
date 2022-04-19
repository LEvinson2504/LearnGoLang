// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func one() {
	defer wg.Done()
	fmt.Println("One")
}
func two() {
	defer wg.Done()
	fmt.Println("Two")
}
func three() {
	defer wg.Done()
	fmt.Println("Three")
}
func main() {
	wg.Add(3)
	go one()
	go two()
	go three()
	fmt.Println("Hello World")
	wg.Wait()
}

