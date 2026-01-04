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
	}
	
	fmt.Println(ConvertToJson(m))
	fmt.Println(ConvertToJson(e))
	fmt.Println(ConvertToJson(t))

}



func ConvertToJson(n any) string{
	b,e := json.Marshal(n)
	if e != nil {
		panic(e)
	}
	
	return  string(b)
}