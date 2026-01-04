package main

import (
	"encoding/json"
	"fmt"
)

type PersonWithOutEmail struct {
	Name string
	Age int
	email string
}
type PersonWithTags struct{
	Name string `json:"FirstName"`
	Age int `json:"Age"`
	Email string `json:"Email"`
	Address Address `json:"Address"`
}

type Address struct{
    City    string 
    Street  string
}


func main() {


	m := map[string]int{
		"one" : 1,
		"two" : 2,
		"three" : 3,
	}
	
	e := PersonWithOutEmail{
		Name: "Matin",
		Age: 20,
		email: "mnasirii829@gmail.com",
	}

	t := PersonWithTags{
		Name: "Mehran",
		Age: 19,
		Email: "Mehran@gmail.com",
		Address: Address{
			City: "Mashhad",
			Street: "Toos",
		},
	}
	
	fmt.Println(CTJMarshal(m))
	fmt.Println(CTJMarshal(e))
	fmt.Println(CTJMarshal(t))
	fmt.Println(CTJMarshalIndent(t))


}



func CTJMarshal(n any) string{
	b,e := json.Marshal(n)
	if e != nil {
		panic(e)
	}
	
	return  string(b)
}

func CTJMarshalIndent(n any) string{
	b,e := json.MarshalIndent(n,"","\t")
	if e != nil {
		panic(e)
	}
	
	return  string(b)
}