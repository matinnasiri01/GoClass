package main

import "fmt"





type Student struct {
   FirstName, LastName string
   Age                 int
   Major               string
}



func main() {


	var s1 = Student{FirstName: "Olivia", Age: 18, }
	s2 := Student{FirstName: "Olivia", LastName: "Martin", Age: 18, Major: "Electrical Engineering"}
	// If we want to initialize a struct without mentioning the names of its fields, all fields must be initialized and the order in which the fields are initialized must be in the order in which they are defined in the struct.
	var s3 = Student{"Sophia", "Smith", 19, "Computer Engineering"}
	fmt.Printf("%+v\n%+v\n%+v\n", s1,s2,s3)



	// Anonymous Struct
	var tmp = struct {
     	ProductName string
     	Price       int
   	}{"Chips", 40000}

   	fmt.Printf("Anonymous Struct: %+v", tmp)


	fmt.Printf("s1 name: %s",s1.FirstName)

	
}

// If all the fields of a struct are comparable, that struct can be defined as a key in a map.
	
type Person struct {
   Name      string
   Age       int
   Favorites []string
}

func maps() {
   people := make([]Person, 0)

   people = append(people, Person{Name: "Emily", Age: 20, Favorites: []string{"running", "watch TV"}})

   p2 := Person{
      Name:      "Joe",
      Age:       30,
      Favorites: []string{"chess"},
   }

   people = append(people, p2)

   for i, v := range people {
      fmt.Printf("%d. Name : %s Age : %d Favorites : ", i+1, v.Name, v.Age)

      for j, f := range v.Favorites {
         fmt.Printf("%d. %s ", j+1, f)
      }

      fmt.Printf("\n")

   }
}