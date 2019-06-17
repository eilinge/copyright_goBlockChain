package main

import "fmt"

// Human ...
type Human struct {
	name string
	age  uint
	sex  uint
}

// SuperMan ...
type SuperMan struct {
	Human
	level uint
}

func (p *Human) setName(n string) {
	p.name = n
}

func (p *Human) setSex(s uint) {
	p.sex = s
}

func (s *SuperMan) setLevel(l uint) {
	s.level = l
}

func main02() {
	h1 := Human{"eilinge", 17, 1}

	var s1 SuperMan
	s1.Human = h1
	s1.setLevel(2)

	s1.setName("duzi") // 直接继承了Human的方法(setName/setSex)
	s1.Human.setName("lin")

	fmt.Printf("the superMan is:%#v", s1)
}
