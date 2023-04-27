package main

import "fmt"

func main() {
	// ➡️ Structs allows to group variables / fields together
	// you define them like this :
	type person struct {
		firstname string
		lastname  string
		age       uint8
	}

	// ➡️ You then initialize them like
	jane := person{
		firstname: "Jane",
		lastname:  "Doe",
		age:       38,
	}
	fmt.Println(jane)

	// ➡️ You can also initialize them without field names but avoid it, it's less clear.
	john := person{
		"John",
		"Doe",
		38,
	}
	fmt.Println(john)

	// ➡️ each field is accessible for reading and writing
	fmt.Println(jane.firstname)
	fmt.Println(jane.lastname)
	fmt.Println(jane.age)
	jane.age += 1

	// ➡️ It's also possible to define Anonymous Structs
	car := struct {
		brand string
		model string
	}{
		"Renault",
		"Clio",
	}
	println(car)

}

// ➡️ You can also attach methods to structs !
type dog struct {
	name string
}

func (d dog) bark() string {
	//This type of Struct method is called a 'Value receiver method'
	// ⚠️ As is, the variable 'd' is read-only, d.name = "newName" won't work
	//To solve this, there is also the 'Pointer receiver method' which we will
	// see after looking the Pointer subject
	return d.name
}

func walkOutTheDog() {
	pet := dog{"pet"}
	pet.bark()
}

// ➡️ Struct usually deserves their own file
// They also have a method like 'NewPerson()' which helps instantiate them
