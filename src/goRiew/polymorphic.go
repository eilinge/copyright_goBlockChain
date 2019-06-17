package main

import (
	"fmt"
)

// Coder ...
type Coder interface {
	coding()
	sleeping()
}

// Pythoner ...
type Pythoner struct {
	name string
	// age  uint8
}

// BlockChainer ...
type BlockChainer struct {
	name string
	// age  uint8
}

func (python *Pythoner) coding() {
	fmt.Printf("this is:%s coding\n", python.name)
}

func (python *Pythoner) sleeping() {
	fmt.Printf("this is:%s sleeping\n", python.name)
}

func (blocker *BlockChainer) coding() {
	fmt.Printf("this is:%s coding\n", blocker.name)
}

func (blocker *BlockChainer) sleeping() {
	fmt.Printf("this is:%s sleeping\n", blocker.name)
}

// Train ...
func Train(name string, num uint8) Coder {
	// var who Coder
	if num == 1 {
		return &Pythoner{name}
	}
	return &BlockChainer{name}

}

func main01() {
	a := Pythoner{"eilinge"}
	b := BlockChainer{"lin"}

	a.coding()
	b.sleeping()

	who := Train("duzi", 1)
	fmt.Println("i am: ", who)
}
