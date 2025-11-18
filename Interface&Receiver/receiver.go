package main

import "fmt"

/// Receiver
type Cat struct{ 
	color string
}

func (c Cat) MakeSound() {

    fmt.Printf("cat color %s say: Meow\n",c.color) 
}
func (c *Cat) ChangeColor(i string) {
	c.color = i
}


func mainReceiver() {

	persianCat := Cat{"White"}
	// know we can use MakeSound() Function.
	persianCat.MakeSound()
	// use pointer to change cat color.
	persianCat.ChangeColor("Black")
	persianCat.MakeSound()


}
