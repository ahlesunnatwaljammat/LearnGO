package main

import (
	"log"
	"os"
)

type Person struct {
	Name string
	Age  int64
}

func init() {
	log.SetOutput(os.Stdout)
}

func setAge(person Person) {
	person.Age = 55
	log.Println("setAge: ", person)
}

func setAgeByRef(person *Person) {
	person.Age = 55
	log.Println("setAgeByRef: ", person)
}

func (person *Person) update() {
	person.Name = "Name is updated by Update method"
}

func main() {

	log.Println("=================== Person by Value ===================")
	person := Person{}
	person.Name = "Noman ali abbasi"
	person.update()
	setAge(person)
	log.Printf("%T, %v", person, person)



	log.Println("=================== Person by Reference ===================")
	person1 := &Person{}
	person1.Name = "Noman ali abbasi"
	setAgeByRef(person1)
	log.Printf("%T, %v", person1, person1)

	log.Println("=================== Person struct as dynamic array ===================")
	people := [...]Person{
		{
			Name: "nabbasi",
			Age:  34,
		}, {
			Name: "fabbasi",
			Age:  33,
		}, {
			Name: "aabbasi",
			Age:  32,
		},
	}

	log.Println(people)

	for i := 0; i < len(people); i++ {
		log.Println(people[i].Name, people[i].Age)
	}

	for i := range people {
		log.Println(people[i].Name, people[i].Age)
	}
}
