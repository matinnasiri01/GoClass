package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
func Some (a,b int) int {
   return  a + b
}

// simple testing pakage:
func TestSome (t *testing.T){
   if Some(2,2) != 4 {
      t.Errorf("Wrong Anser!")
   }
}

// use https://github.com/stretchr/testify lib:
func TestSomeTestify (t *testing.T){
   // assert equality
  assert.Equal(t, 123, 123, "they should be equal")

  // assert inequality
  assert.NotEqual(t, 123, 456, "they should not be equal")
  



  
}


func main() { }