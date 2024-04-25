package main

import "fmt"

type Creature struct {
	Species string
}

func (c *Creature) Reset() { // receptor pointer
	c.Species = ""
}

func main() {
	//var creature string = "shark"
	//var pointer *string = &creature

	//fmt.Println("creature =", creature) // shark
	//fmt.Println("pointer =", pointer)   // memory address
	//fmt.Println("*pointer =", *pointer) // this is known as "dereferencing" or "indirecting"

	//*pointer = "jellyfish" // change value
	//fmt.Println("*pointer =", *pointer)

	//fmt.Println("creature =", creature) // after change value

	//var creature *Creature
	//creature = &Creature{Species: "shark"}
	//fmt.Printf("1) %+v\n", creature)
	//changeCreature(creature)
	//fmt.Printf("3) %+v\n", creature)

	var creature Creature = Creature{Species: "shark"}

	fmt.Printf("1) %+v\n", creature)
	creature.Reset()
	fmt.Printf("2) %+v\n", creature)
}

func changeCreature(creature *Creature) {
	if creature == nil {
		fmt.Println("creature is nil")
		return
	}

	creature.Species = "jellyfish"
	fmt.Printf("2) %+v\n", creature)
}
