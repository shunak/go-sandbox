package main

import (
	"fmt"
)

type Person struct {
	namae string
	toshi uint
}

// pointer pass method
func (p *Person) renamae1(newnamae string) {
	p.namae = newnamae
}

// value pass method
func (p Person) renamae2(newnamae string) {
	p.namae = newnamae
}

func main() {

	fmt.Println("Person.renamae1->%T\n", (*Person).renamae1)
	fmt.Println("Person.renamae2->%T\n", (Person).renamae2)

	personA := &Person{namae: "Atsuko", toshi: 19}
	oldnamae := personA.namae

	personA.renamae1("Naoko")

	fmt.Println("renamae1 exec renamaed. %s -> %s\n", oldnamae, personA.namae)

	personB := &Person{namae: "Tomoka", toshi: 19}
	oldnamae = personB.namae

	personB.renamae2("Hanako")

	fmt.Println("renamae2 exec renamaed. %s -> %s\n", oldnamae, personB.namae)

}
